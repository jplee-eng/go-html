package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonpatricklee/gowebserver/handlers"
)

func Routes(r *gin.Engine) {
	api := r.Group("/api")
	// health status routes
	api.GET("/health", handlers.HealthHandler)
	api.GET("/isActive", handlers.HealthHandler)
	api.GET("/active", handlers.HealthHandler)
	api.GET("/ping", handlers.HealthHandler)
}
