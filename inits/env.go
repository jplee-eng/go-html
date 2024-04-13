package inits

import (
	"log"

	"github.com/joho/godotenv"
)

func ReadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load the .env file!")
	}
}
