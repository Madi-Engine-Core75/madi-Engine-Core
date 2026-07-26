package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api/v1/auth/login", handleLogin)
	http.ListenAndServe(":8080", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	res := Response{
		Status:  "success",
		Message: "Authentication gateway ready",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
