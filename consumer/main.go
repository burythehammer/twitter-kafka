package main

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"sync"
)

const topicPopularTweets = "popular_tweets"
const topicAllTweets = "twitter_tweets"

var consumerConfig = &kafka.ConfigMap{
	"bootstrap.servers": "localhost",
	"group.id":          "myGroup",
	"auto.offset.reset": "earliest",
}

func main() {
	var wg sync.WaitGroup

	go consumeTweets(topicPopularTweets, &wg)
	wg.Add(1)

	go consumeTweets(topicAllTweets, &wg)
	wg.Add(1)

	allTweets := KafkaConsumer{
		config: consumerConfig,
		topic:  topicPopularTweets,
	}
	go allTweets.consume()

	wg.Wait()
}

func consumeTweets(topic string, wg *sync.WaitGroup) {
	popularTweets := KafkaConsumer{
		config: consumerConfig,
		topic:  topic,
	}
	popularTweets.consume()
	wg.Done()
}
