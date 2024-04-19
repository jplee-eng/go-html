package main

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/jasonpatricklee/gowebserver/inits"

	"github.com/jasonpatricklee/gowebserver/internal/routes"
)

func init() {
	inits.ReadEnv()
	inits.StartSentry()
}
func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	routes.Routes(r)
	r.Run(":3040")
}
