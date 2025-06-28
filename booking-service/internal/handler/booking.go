package handler

import (
	"booking-service/internal/kafka"
	"booking-service/internal/model"
	"encoding/json"
	"log"
	"net/http"
)

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	var req model.BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if req.Token == "" || req.PassengerName == "" || req.TrainID == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	jsonReq, _ := json.Marshal(req)
	kafka.ProduceBooking(string(jsonReq))

	log.Printf("Queued booking for %s on train %s", req.PassengerName, req.TrainID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "queued"})
}
