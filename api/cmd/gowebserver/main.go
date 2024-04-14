package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/jasonpatricklee/gowebserver/inits"
	"github.com/jasonpatricklee/gowebserver/internal/routes"
)

func init() {
	// inits.ReadEnv()
	fmt.Print("Read inits after editing env db string")
}
func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
