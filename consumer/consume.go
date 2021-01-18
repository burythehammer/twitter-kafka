package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
)

type KafkaConsumer struct {
	config *kafka.ConfigMap
	topic  string
}

const fileName = "data/tweets.json"

func (dkc KafkaConsumer) consume() {
	c, err := kafka.NewConsumer(dkc.config)

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{dkc.topic}, nil)

	if err != nil {
		panic(err)
	}

	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Received text: %+v", string(msg.Value))
			appendToFile(msg.Value)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func appendToFile(text []byte) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(text); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
