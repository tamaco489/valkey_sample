package router

import (
	"net/http"

	"github.com/tamaco489/valkey_sample/internal/handler"
)

// Setup configures the router with all routes
func Setup() *http.ServeMux {
	mux := http.NewServeMux()

	// API v1 group
	mux.HandleFunc("GET /api/v1/health", handler.HealthCheck)

	return mux
}
