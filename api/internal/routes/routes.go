package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonpatricklee/gowebserver/handlers"
)

func Routes(r *gin.Engine) {
	// health status routes
	r.GET("/health", handlers.HealthHandler)
	r.GET("/isActive", handlers.HealthHandler)
	r.GET("/active", handlers.HealthHandler)
	r.GET("/ping", handlers.HealthHandler)
}
