package main

import (
	"go-html/inits"
	"go-html/internal"
	"log"
	"net/http"
)

func init() {
	inits.ReadEnv()
}

func main() {
	mux := http.NewServeMux()
	internal.SetupRoutes(mux)

	server := &http.Server{
		Addr:    ":3040",
		Handler: mux,
	}

	log.Printf("Server starting on port 3040")
	log.Fatal(server.ListenAndServe())
}
