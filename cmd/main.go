package main

import (
	"fmt"
	"net/http"

	"logistics-app/internal/handlers"
)

func main() {
	// Routes
	http.HandleFunc("/shipments", handlers.GetShipments)
	http.HandleFunc("/create", handlers.CreateShipment)
	http.HandleFunc("/health", handlers.HealthCheck)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}