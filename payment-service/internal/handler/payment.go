package handler

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"payment-service/internal/model"
	"time"
)

func PayHandler(w http.ResponseWriter, r *http.Request) {
	var req model.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payment request", http.StatusBadRequest)
		return
	}

	// Simulate payment delay
	time.Sleep(500 * time.Millisecond)

	txnID := "TXN-" + req.Reference + "-" + randomDigits(6)
	resp := model.PaymentResponse{
		Status: "success",
		TxnId:  txnID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	log.Printf("Payment completed for %s: ₹%.2f -> %s\n", req.Reference, req.Amount, txnID)
}

func randomDigits(n int) string {
	const digits = "0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = digits[rand.Intn(len(digits))]
	}
	return string(b)
}
