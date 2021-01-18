package main

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const topicPopularTweets = "popular_tweets"

var consumerConfig = &kafka.ConfigMap{
	"bootstrap.servers": "localhost",
	"group.id":          "myGroup",
	"auto.offset.reset": "earliest",
}

func main() {
	popularTweets := KafkaConsumer{
		config: consumerConfig,
		topic:  topicPopularTweets,
	}
	popularTweets.consume()
}
