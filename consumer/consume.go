package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
	"time"
)

type KafkaConsumer struct {
	config *kafka.ConfigMap
	topic  string
}

func (kc KafkaConsumer) consume() {
	c, err := kafka.NewConsumer(kc.config)

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{kc.topic}, nil)

	if err != nil {
		panic(err)
	}

	defer c.Close()
	folder := createTopicFolder(kc.topic)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %+v", msg)
			createTweetFile(msg.Value, folder, msg.String())
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func createTopicFolder(topic string) string {
	now := time.Now()
	directoryName := fmt.Sprintf("./data/%d/%s", now.Unix(), topic)
	err := os.MkdirAll(directoryName, 0755)

	if err != nil {
		err := fmt.Errorf("could not create data directory for topic %s, error: %s", topic, err.Error())
		panic(err)
	}

	return directoryName
}

func createTweetFile(text []byte, dir string, fileName string) {
	// If the file doesn't exist, create it, or append to the file

	filePath := fmt.Sprintf("%s/%s.json", dir, fileName)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
