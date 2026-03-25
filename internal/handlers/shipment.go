package handlers

import (
	"encoding/json"
	"net/http"
)

// Shipment structure
type Shipment struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Location string `json:"location"`
}

// In-memory storage (for demo)
var shipments = []Shipment{
	{ID: "1", Status: "In Transit", Location: "Delhi"},
	{ID: "2", Status: "Delivered", Location: "Mumbai"},
}

// GET /shipments
func GetShipments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shipments)
}

// POST /shipments
func CreateShipment(w http.ResponseWriter, r *http.Request) {
	var newShipment Shipment

	err := json.NewDecoder(r.Body).Decode(&newShipment)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shipments = append(shipments, newShipment)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newShipment)
}

// GET /health
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}