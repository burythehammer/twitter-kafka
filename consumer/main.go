package main

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

const TwitterTopic = "twitter-tweets"

var consumerConfig = &kafka.ConfigMap{
	"bootstrap.servers": "localhost",
	"group.id":          "myGroup",
	"auto.offset.reset": "earliest",
}

func main() {
	dkc := KafkaConsumer{
		config: consumerConfig,
		topic:  TwitterTopic,
	}
	dkc.consume()
}
