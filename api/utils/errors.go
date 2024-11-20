package utils

import (
	"log"
)

func HandleError(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}

func HandleFatalError(err error) {
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}
