package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/antoinecrochet/transport-rennes-api/internal/core/port"
	"github.com/gorilla/mux"
)

type Application struct {
	search port.Search
	port   string
}

// Application constructor
func NewApplication(search port.Search) *Application {
	return &Application{
		search: search,
		port:   "8080",
	}
}

// Start application
func (app *Application) Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/search/upcomingbus", app.searchUpcomingBus).Methods("POST")

	log.Printf("Starting application on port %s ...", app.port)
	log.Fatal(http.ListenAndServe(":"+app.port, router))
}

// Get upcoming bus
func (app *Application) searchUpcomingBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	var data SearchBus
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	// only the stop name is mandatory
	if data.Stop == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Error{"stop-field-missing"})
		return
	}

	result, err := app.search.SearchUpcomingBus(data.BusLine, data.Stop, data.Destination)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Error{err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	// If no bus left
	if result.Count == 0 {
		json.NewEncoder(w).Encode(&SearchUpcomingBusResponse{Message: "Aucun bus disponible", Count: result.Count, Hits: []HitResponse{}})
		return
	}

	// Convert domain model into api model
	hits := make([]HitResponse, len(result.Hits))
	for i, element := range result.Hits {
		hits[i] = HitResponse{
			BusLineName: element.BusLineName,
			BusStopName: element.BusStopName,
			Destination: element.Destination,
			Departure:   element.Departure,
		}
	}

	// Generate message
	var message string
	if result.Count >= 2 {
		message = fmt.Sprintf("Prochain bus dans %d min, le suivant dans %d min",
			getDelay(&result.Hits[0].Departure),
			getDelay(&result.Hits[1].Departure))
	} else if result.Count == 1 {
		message = fmt.Sprintf("Prochain bus dans %d", getDelay(&result.Hits[0].Departure))
	}

	json.NewEncoder(w).Encode(&SearchUpcomingBusResponse{Message: message, Count: result.Count, Hits: hits})
}

// Return delay before departure in minutes
func getDelay(departure *time.Time) int {
	return int(departure.Sub(time.Now().UTC()).Minutes())
}
