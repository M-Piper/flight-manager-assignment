package main

import (
	"testing"
	"time"
)

// Helper function to reset the FlightSchedule for testing.
func resetFlightSchedule() {
	FlightSchedule = []Flight{}
}

// Test adding a new flight correctly.
func TestAddFlight(t *testing.T) {
	resetFlightSchedule() // Resetting the flight schedule before the test.

	flight := Flight{
		FlightNumber: "AC 301",
		Origin:       "YVR",
		Destination:  "YYZ",
		Departure:    time.Date(2024, 9, 20, 15, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 20, 22, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}

	AddFlight(flight) // Add the flight.
	if len(FlightSchedule) != 1 {
		t.Errorf("Expected 1 flight in the schedule, got %d", len(FlightSchedule))
	}
}

// Test editing a flight correctly.
func TestEditFlight(t *testing.T) {
	resetFlightSchedule()

	flight := Flight{
		FlightNumber: "AC 401",
		Origin:       "YVR",
		Destination:  "YYZ",
		Departure:    time.Date(2024, 9, 21, 15, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 21, 22, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}
	AddFlight(flight)

	newFlight := Flight{
		FlightNumber: "AC 402",
		Origin:       "YYZ",
		Destination:  "YVR",
		Departure:    time.Date(2024, 9, 21, 16, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 21, 23, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}
	EditFlight("AC 401", newFlight)

	if FlightSchedule[0].FlightNumber != "AC 402" {
		t.Errorf("Expected flight number to be 'AC 402', got '%s'", FlightSchedule[0].FlightNumber)
	}
}

// Test trying to edit a flight that does not exist.
func TestEditFlightNotFound(t *testing.T) {
	resetFlightSchedule()

	flight := Flight{
		FlightNumber: "AC 501",
		Origin:       "YVR",
		Destination:  "YYZ",
		Departure:    time.Date(2024, 9, 22, 15, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 22, 22, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}
	AddFlight(flight)

	newFlight := Flight{
		FlightNumber: "AC 502",
		Origin:       "YYZ",
		Destination:  "YVR",
		Departure:    time.Date(2024, 9, 22, 16, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 22, 23, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}
	EditFlight("AC 999", newFlight) // This flight does not exist

	if len(FlightSchedule) != 1 {
		t.Errorf("Expected 1 flight in the schedule, got %d", len(FlightSchedule))
	}
}

// Test deleting a flight successfully.
func TestDeleteFlight(t *testing.T) {
	resetFlightSchedule()

	flight := Flight{
		FlightNumber: "AC 601",
		Origin:       "YVR",
		Destination:  "YYZ",
		Departure:    time.Date(2024, 9, 23, 15, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 23, 22, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}
	AddFlight(flight)

	DeleteFlight("AC 601") // Attempt to delete the flight.

	if len(FlightSchedule) != 0 {
		t.Errorf("Expected 0 flights in the schedule, got %d", len(FlightSchedule))
	}
}

// Test deleting a flight with incorrect syntax.
func TestDeleteFlightIncorrectSyntax(t *testing.T) {
	resetFlightSchedule()

	// Simulate incorrect input
	DeleteFlight("AC 999") // Attempt to delete a flight that does not exist.
	if len(FlightSchedule) != 0 {
		t.Errorf("Expected 0 flights in the schedule, got %d", len(FlightSchedule))
	}
}

// Test searching for a flight that exists.
func TestSearchFlightExists(t *testing.T) {
	resetFlightSchedule()

	flight := Flight{
		FlightNumber: "AC 701",
		Origin:       "YVR",
		Destination:  "YYZ",
		Departure:    time.Date(2024, 9, 24, 15, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 24, 22, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}
	AddFlight(flight)

	// Capture output and check if the flight is found (you can use a testing library to capture stdout if needed).
	SearchFlights("Sep 24 2024", "YVR", "YYZ") // This should find the flight.
}

// Test searching for a flight that does not exist.
func TestSearchFlightNotFound(t *testing.T) {
	resetFlightSchedule()

	flight := Flight{
		FlightNumber: "AC 801",
		Origin:       "YVR",
		Destination:  "YYZ",
		Departure:    time.Date(2024, 9, 25, 15, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 25, 22, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}
	AddFlight(flight)

	// Capture output and check if the flight is found (you can use a testing library to capture stdout if needed).
	SearchFlights("Sep 26 2024", "", "") // This should not find the flight.
}
