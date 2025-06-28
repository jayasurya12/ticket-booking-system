package main

import (
	"booking-service/internal/handler"
	"booking-service/internal/kafka"
	"log"
	"net/http"
)

func main() {
	kafka.InitProducer()
	go kafka.StartConsumer()

	http.HandleFunc("/book", handler.BookingHandler)

	log.Println("Booking Service running on :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
