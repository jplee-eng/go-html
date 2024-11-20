package inits

import (
	"go-html/utils"

	"github.com/joho/godotenv"
)

func ReadEnv() {
	err := godotenv.Load()
	utils.HandleError(err)
}
