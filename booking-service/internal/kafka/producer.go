package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var producer *kafka.Producer

func InitProducer() {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Update if you're using Kubernetes
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
}

func ProduceBooking(msg string) {
	topic := "bookings"
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}, nil)

	if err != nil {
		log.Println("Failed to send message to Kafka:", err)
	}
}
