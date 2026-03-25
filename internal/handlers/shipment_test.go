package handlers

import "testing"

// Test if shipments list is not empty
func TestGetShipments(t *testing.T) {
	if len(shipments) == 0 {
		t.Error("Expected shipments, but got empty list")
	}
}

// Test adding a new shipment
func TestCreateShipment(t *testing.T) {
	initialCount := len(shipments)

	newShipment := Shipment{
		ID:       "99",
		Status:   "Testing",
		Location: "TestCity",
	}

	shipments = append(shipments, newShipment)

	if len(shipments) != initialCount+1 {
		t.Error("Shipment was not added correctly")
	}
}