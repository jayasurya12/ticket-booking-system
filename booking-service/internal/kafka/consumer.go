package kafka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"booking-service/internal/model"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func StartConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Use service name in K8s
		"group.id":          "booking-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal("Kafka consumer creation error:", err)
	}

	err = c.SubscribeTopics([]string{"bookings"}, nil)
	if err != nil {
		log.Fatal("Kafka topic subscription error:", err)
	}

	log.Println("Kafka consumer started...")

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading Kafka message:", err)
			continue
		}

		var booking model.BookingRequest
		if err := json.Unmarshal(msg.Value, &booking); err != nil {
			log.Println("Failed to unmarshal booking:", err)
			continue
		}

		// Simulate payment
		paymentReq := map[string]interface{}{
			"token":     booking.Token,
			"amount":    500,
			"reference": booking.TrainID + "_" + booking.PassengerName,
		}
		payload, _ := json.Marshal(paymentReq)

		resp, err := http.Post("http://localhost:8083/pay", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			log.Println("Payment call failed:", err)
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Booking confirmed for %s | Payment: %s\n", booking.PassengerName, string(body))
	}
}
