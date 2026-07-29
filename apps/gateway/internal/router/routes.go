package router

import (
"net/http"

"gateway/internal/handler"
)

func SetupRoutes() *http.ServeMux {
mux := http.NewServeMux()

// Register handlers
mux.HandleFunc("/health", handler.HealthCheck)

return mux
}
