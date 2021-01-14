package main

import (
	"github.com/burythehammer/twitter-kafka/producer/twitter"
	"github.com/dghubble/go-twitter/twitter"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
	"strings"
)

const TwitterTopic = "twitter-tweets"

var producerConfig = kafka.ConfigMap{
	"bootstrap.servers":  "localhost",
	"enable.idempotence": "true",
	"compression.type":   "snappy",
	"linger.ms":          "20",
}

func main() {
	dkp := KafkaProducer{
		config: producerConfig,
		topic:  TwitterTopic,
	}

	searchArgs := os.Args[1:]

	if len(searchArgs) == 0 {
		panic("Requires at least one search term")
	}

	query := strings.Join(searchArgs, " OR ")

	client := twitterclient.CreateTwitterClient()

	tweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
	})

	if err != nil {
		panic(err)
	}

	tweetStatuses := getTweetText(tweets.Statuses)
	dkp.Produce(tweetStatuses)
}

func getTweetText(tweets []twitter.Tweet) []string {

	var text []string

	for _, tweet := range tweets {
		text = append(text, tweet.Text)
	}

	return text
}
