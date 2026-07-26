package router

import (
	"net/http"
	"github.com/MadiEngine-Core75/Madi-Engine-Core/apps/gateway/internal/handler"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// نقطة نهاية المصادقة المطلوبة
	mux.HandleFunc("/api/v1/auth/login", handler.HandleLogin)

	return mux
}
