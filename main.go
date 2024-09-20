package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

// Defining individual flights
type Flight struct {
	FlightNumber string
	Origin       string
	Destination  string
	Departure    time.Time
	Arrival      time.Time
}

// Creating slice to hold lists of the Flight struct (using this in-memory list in lieu of real database)
var FlightSchedule []Flight

// Function allowing user to add flight in the cli
func AddFlight(flight Flight) error {
	if flight.FlightNumber == "" || flight.Origin == "" || flight.Destination == "" {
		return fmt.Errorf("invalid flight data")
	}
	FlightSchedule = append(FlightSchedule, flight) // allows user to add flight to FlightSchedule slice.
	fmt.Println("Flight added successfully.")
	return nil
}

// Function to search for flights by date, origin, and destination
func SearchFlights(date string, origin string, destination string) {
	layout := "Jan 02 2006"                     // defining format/layout of date (allows go to correctly identify month, day, year)
	searchDate, err := time.Parse(layout, date) // parsing date
	if err != nil {                             // error handling - occurs if the date provided is not in the correct format (ie day before month)
		fmt.Println("Date format incorrect - please use 'mmm dd yyyy' (example: 'Sep 18 2024')")
		return
	}
	fmt.Printf("You have asked to search for flights on %s from %s to %s\n", date, origin, destination) // returning user search parameters

	// walk through the FlightSchedule slice to find matching flight instances
	for _, flight := range FlightSchedule {
		// if statement checks origin, destiation and date (all must match parameters identified above)
		if (origin == "" || flight.Origin == origin) &&
			(destination == "" || flight.Destination == destination) &&
			flight.Departure.Format(layout) == searchDate.Format(layout) {

			// prints results
			fmt.Printf("Flight Number: %s, Origin: %s, Destination: %s, Departure: %s, Arrival: %s\n",
				flight.FlightNumber, flight.Origin, flight.Destination,
				flight.Departure.Format("Jan 02, 2006 at 15:04 MST"), // adjusting formatting of date and time for display
				flight.Arrival.Format("Jan 02, 2006 at 15:04 MST"))   // adjusting formatting of date and time for display
		}
	}
}

// Function to delete flight from FlightSchedule slice (identified by flight number)
func DeleteFlight(flightNumber string) {
	for i, flight := range FlightSchedule { // walk through the entire slice
		if flight.FlightNumber == flightNumber { // if statement (to remove flight instance if flight number matches)
			FlightSchedule = append(FlightSchedule[:i], FlightSchedule[i+1:]...) // removes flight
			fmt.Println("Flight deleted successfully.")
			return
		}
	}
	fmt.Println("That flight number does not correspond to any currently scheduled flights.") // Return for no match found
}

// EditFlight edits an existing flight.
func EditFlight(flightNumber string, newFlight Flight) {
	for i, flight := range FlightSchedule { // Iterates through the FlightSchedule.
		if flight.FlightNumber == flightNumber { // If the flight number matches, update the flight.
			FlightSchedule[i] = newFlight               // Replaces the old flight with the new flight.
			fmt.Println("Flight updated successfully.") // Prints confirmation message.
			return
		}
	}
	fmt.Println("Flight not found.") // If no matching flight is found, print an error message.
}

func main() {
	// Some test data to populate the slice for the purpose of the exercise - would be replaced by database
	flight1 := Flight{
		FlightNumber: "AC 101",
		Origin:       "YVR",
		Destination:  "YYZ",
		Departure:    time.Date(2024, 9, 18, 15, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 18, 22, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}
	flight2 := Flight{
		FlightNumber: "AC 102",
		Origin:       "YYC",
		Destination:  "YUL",
		Departure:    time.Date(2024, 9, 18, 14, 0, 0, 0, time.FixedZone("MST", -7*3600)),
		Arrival:      time.Date(2024, 9, 18, 19, 30, 0, 0, time.FixedZone("EST", -5*3600)),
	}
	flight3 := Flight{
		FlightNumber: "WS 203",
		Origin:       "YEG",
		Destination:  "YVR",
		Departure:    time.Date(2024, 9, 18, 16, 30, 0, 0, time.FixedZone("MST", -7*3600)),
		Arrival:      time.Date(2024, 9, 18, 17, 30, 0, 0, time.FixedZone("PST", -8*3600)),
	}
	flight4 := Flight{
		FlightNumber: "WS 204",
		Origin:       "YVR",
		Destination:  "YYC",
		Departure:    time.Date(2024, 9, 18, 18, 0, 0, 0, time.FixedZone("PST", -8*3600)),
		Arrival:      time.Date(2024, 9, 18, 21, 0, 0, 0, time.FixedZone("MST", -7*3600)),
	}

	FlightSchedule = append(FlightSchedule, flight1, flight2, flight3, flight4) // appends the data above to the FlightSchedule slice

	// Define flags for user input from the command line
	addFlag := flag.String("add", "", "Add a new flight with format 'FlightNumber,Origin,Destination,Departure,Arrival'")
	searchDate := flag.String("search-date", "", "Search flights by date in format 'Sep 18, 2024'")
	origin := flag.String("origin", "", "Search flights by origin")
	destination := flag.String("destination", "", "Search flights by destination")
	editFlag := flag.String("edit", "", "Edit an existing flight with format 'FlightNumber,NewFlightNumber,NewOrigin,NewDestination,NewDeparture,NewArrival'")
	deleteFlag := flag.String("delete", "", "Delete a flight by FlightNumber")

	flag.Parse() // Parses command-line arguments

	if *addFlag != "" { // If the add flag is provided, flight is added to FlightSchedule slice
		parts := strings.Split(*addFlag, ",") // Splits the input string by commas
		if len(parts) != 5 {                  // Confirming the input contains the right number of inputs
			fmt.Println("Something went wrong, be sure to enter five pieces of information separated by commas using the order 'FlightNumber,Origin,Destination,Departure,Arrival'")
			return
		}
		// Parsing departure and arrival times
		departure, _ := time.Parse("Jan 02 2006 15:04 MST", parts[3])
		arrival, _ := time.Parse("Jan 02 2006 15:04 MST", parts[4])
		flight := Flight{ // Creating new flight object
			FlightNumber: parts[0],
			Origin:       parts[1],
			Destination:  parts[2],
			Departure:    departure,
			Arrival:      arrival,
		}
		AddFlight(flight) // Adds the new flight to the schedule.
	}

	if *searchDate != "" { // If the search-date flag is provided, search for flights
		SearchFlights(*searchDate, *origin, *destination) // Calls the search function with the provided parameters.
	}

	if *editFlag != "" { // If the edit flag is provided, edit an existing flight.
		parts := strings.Split(*editFlag, ",")
		if len(parts) != 6 {
			fmt.Println("Invalid entry - enter the flight number you wish to alter followed by the revised input (example: AC 101, AC 402, YYC, YUL, Sep 18 2024 at 14:00 MST, SEP 18 2024 at 19:30 EST)")
			return
		}
		// Parses the new departure and arrival times.
		departure, _ := time.Parse("Jan 02, 2006 15:04 MST", parts[4])
		arrival, _ := time.Parse("Jan 02, 2006 15:04 MST", parts[5])
		newFlight := Flight{ // Creates a new flight object.
			FlightNumber: parts[1],
			Origin:       parts[2],
			Destination:  parts[3],
			Departure:    departure,
			Arrival:      arrival,
		}
		EditFlight(parts[0], newFlight) // Calls the edit function with the provided parameters.
	}

	if *deleteFlag != "" { // If the delete flag is provided, delete the flight.
		DeleteFlight(*deleteFlag) // Calls the delete function with the flight number.
	}
}
