package main

import (
"log"
"net/http"

"gateway/internal/router"
)

func main() {
mux := router.SetupRoutes()

port := ":8080"
log.Printf("Gateway is running on port %s...", port)
if err := http.ListenAndServe(port, mux); err != nil {
log.Fatalf("Server failed to start: %v", err)
}
}
