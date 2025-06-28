package handler

import (
	"encoding/json"
	"log"
	"login-service/internal/model"
	"net/http"

	"github.com/google/uuid"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Username == "" || req.Password == "" {
		http.Error(w, "Invalid login request", http.StatusBadRequest)
		return
	}

	// Token generation using UUID (dummy)
	token := uuid.New().String()

	res := model.LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	log.Printf("User '%s' logged in. Token: %s\n", req.Username, token)
}
