package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"sport_classes/handlers"
)

func main() {
	// Print current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory: ", err)
	}
	log.Println("Current working directory:", wd)
	// Retrieve the port from the environment variables
	err = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in .env file")
	}

	// Define the HTTP routes
	http.HandleFunc("/parse", handlers.Parse)

	// Start the HTTP server
	log.Printf("Server is running on port %s...\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
