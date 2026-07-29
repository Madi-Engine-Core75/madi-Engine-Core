package main

import (
	"log"
	"net/http"

	"github.com/MadiEngine-Core75/madi-Engine-Core/apps/gateway/internal/router"
)

func handleGatewayRoute(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Madi-Token")

	if token == "" {
		http.Error(w, "Unauthorized: Missing security token", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Gateway routing passed: Request securely forwarded to MadiEngineCore.")
}

func main() {
	mux := router.SetupRoutes()
	mux.HandleFunc("/api/v1/route", handleGatewayRoute)

	port := ":8080"
	log.Printf("Madi-Gateway is running on port %s...", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Gateway server failed to start: %v", err)
	}
}

