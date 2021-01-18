package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaProducer struct {
	config kafka.ConfigMap
	topic  string
}

func (kp KafkaProducer) Produce(messages []string) {
	p, err := kafka.NewProducer(&producerConfig)
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	for _, word := range messages {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
