package inits

import (
	"github.com/jasonpatricklee/gowebserver/utils"
	"github.com/joho/godotenv"
)

func ReadEnv() {
	err := godotenv.Load()
	utils.HandleError(err)
}
