package awesomeProject

import (
	"log"
	"net/http"
	"os"
	"sport_classes/handlers"
)

func main() {

	// Retrieve the port from the environment variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in .env file")
	}

	// Define the HTTP routes
	http.HandleFunc("/parse", handlers.Parse)

	// Start the HTTP server
	log.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
