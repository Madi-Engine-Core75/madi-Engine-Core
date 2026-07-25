package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status    string    `json:"status"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func handleAccountAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	var authReq AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid login payload"})
		return
	}

	res := AuthResponse{
		Status:    "authenticated",
		Token:     "madi_secure_vault_token_sig",
		ExpiresAt: time.Now().Add(time.Hour * 2),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

