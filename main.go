package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Setting up the environment
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the env")
	}

	// Set up an HTTP server and listen on the given port
	router := chi.NewRouter()
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v\n", portString)

	// Start the server, if any error appear immediately stop the program
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
