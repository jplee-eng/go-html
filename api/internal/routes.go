package internal

import (
	"go-html/internal/handlers"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", handlers.HealthHandler)
}
