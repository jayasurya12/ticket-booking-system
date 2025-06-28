package main

import (
	"log"
	"login-service/internal/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handler.LoginHandler)

	log.Println("Login Service running on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
