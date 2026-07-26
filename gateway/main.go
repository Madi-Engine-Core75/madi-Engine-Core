package main

import (
	"encoding/json"
	"net/http"
	"log"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status string `json:"status"`
	Token  string `json:"token,omitempty"`
	Error  string `json:"error,omitempty"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w.Encode(AuthResponse{Status: "error", Error: "Invalid request payload"}))
		return
	}

	// تحقق مبدئي من بيانات الاعتماد (يتم ربطه لاحقاً مع النواة والـ Vault)
	w.Header().Set("Content-Type", "application/json")
	if req.Username == "admin" && req.Password == "secure_password" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w.Encode(AuthResponse{Status: "success", Token: "madi_core_jwt_token_sample"}))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w.Encode(AuthResponse{Status: "unauthorized", Error: "Invalid credentials"}))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/auth/login", handleLogin)

	log.Println("Go Gateway running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
