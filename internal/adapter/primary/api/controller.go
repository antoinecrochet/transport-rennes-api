package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/antoinecrochet/transport-rennes-api/internal/core/model"
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
	router.HandleFunc("/upcomingbus", app.getUpcomingBus).Methods("GET")

	log.Printf("Starting application on port %s ...", app.port)
	log.Fatal(http.ListenAndServe(":"+app.port, router))
}

// Get upcoming bus
func (app *Application) getUpcomingBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	var data SearchBus
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	// only the stop name is mandatory
	if data.Stop == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{"Aucun bus disponible"})
		return
	}

	result, _ := app.search.SearchUpcomingBus(data.BusLine, data.Stop, data.Destination)

	w.WriteHeader(http.StatusOK)
	// If no bus left
	if result.Count == 0 {
		json.NewEncoder(w).Encode(Message{"Aucun bus disponible"})
		return
	}

	// store records by destionation
	x := make(map[string][]model.PublicTransport)
	for _, record := range result.Hits {
		x[record.Destination] = append(x[record.Destination], record)
	}

	// TODO: give next departure for each key of x map (destination)
	// Generate message
	message := Message{}
	if result.Count >= 2 {
		message.Message = fmt.Sprintf("Prochain bus dans %d min, le suivant dans %d min",
			getDelay(&result.Hits[0].Departure),
			getDelay(&result.Hits[1].Departure))
	} else if result.Count == 1 {
		message.Message = fmt.Sprintf("Prochain bus dans %d", getDelay(&result.Hits[0].Departure))
	}

	json.NewEncoder(w).Encode(message)
}

// Return delay before departure in minutes
func getDelay(departure *time.Time) int {
	return int(departure.Sub(time.Now().UTC()).Minutes())
}
