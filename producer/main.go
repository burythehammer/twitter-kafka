package main

import (
	"encoding/json"
	"fmt"
	"github.com/burythehammer/twitter-kafka/producer/twitter"
	"github.com/dghubble/go-twitter/twitter"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
	"strings"
)

const TwitterTopic = "twitter_tweets"

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

	tweetStatuses := marshalTweet(tweets.Statuses)
	dkp.Produce(tweetStatuses)
}

func marshalTweet(tweets []twitter.Tweet) []string {

	var text []string

	for _, tweet := range tweets {

		marshal, err := json.Marshal(tweet)

		if err != nil {
			err := fmt.Errorf("could not marshal tweet %+v, error: %s", tweet, err.Error())
			panic(err)
		}

		text = append(text, string(marshal))
	}

	return text
}
