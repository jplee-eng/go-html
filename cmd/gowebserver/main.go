package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonpatricklee/gowebserver/inits"
	"github.com/jasonpatricklee/gowebserver/internal/routes"
)

func init() {
	inits.ReadEnv()
}
func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
