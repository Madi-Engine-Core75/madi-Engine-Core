package main

import (
	"log"
	"net/http"
	"github.com/MadiEngine-Core75/Madi-Engine-Core/apps/gateway/internal/router"
)

func main() {
	mux := router.SetupRoutes()

	port := ":8080"
	log.Printf("Gateway is running on port %s...", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
