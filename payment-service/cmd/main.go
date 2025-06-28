package main

import (
	"log"
	"net/http"
	"payment-service/internal/handler"
)

func main() {
	http.HandleFunc("/pay", handler.PayHandler)

	log.Println("Payment Service running on :8083")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
}
