package router

import (
	"net/http"

	"github.com/MadiEngine-Core75/madi-Engine-Core/apps/gateway/internal/handler"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/health", handler.HealthCheck)

	return mux
}

