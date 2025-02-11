package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/antoinecrochet/transport-rennes-api/opendatasoft"
	"github.com/gorilla/mux"
)

type Application struct {
	odsClient opendatasoft.OpendatasoftClient
	port      string
}

// Application constructor
func New(configurationFile string) *Application {
	config := opendatasoft.ReadConfigFile(configurationFile)
	return &Application{
		odsClient: *opendatasoft.New(*config),
		port:      "8080",
	}
}

// Start application
func (app *Application) Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/upcomingbus", app.getUpcomingBus).Methods("GET")

	log.Printf("Starting application on port %s ...", app.port)
	log.Fatal(http.ListenAndServe(":"+app.port, router))
}

// Get one apartment by ID
func (app *Application) getUpcomingBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	var data SearchBus
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	// only the stop name is mandatory
	if data.Stop == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	upcomingBus := app.odsClient.GetUpcomingBus(data.BusLine, data.Stop, data.Destination)
	// If no bus left
	if upcomingBus.NHits == 0 {
		json.NewEncoder(w).Encode(Message{"Aucun bus disponible"})
		return
	}

	// sort records by departure time
	sort.SliceStable(upcomingBus.Records, func(i, j int) bool {
		return upcomingBus.Records[i].Information.Departure.Before(upcomingBus.Records[j].Information.Departure)
	})

	// store records by destionation
	x := make(map[string][]opendatasoft.UpcomingBusRecord)
	for _, record := range upcomingBus.Records {
		x[record.Information.Destination] = append(x[record.Information.Destination], record)
	}

	// TODO: give next departure for each key of x map (destination)
	// Generate message
	message := Message{}
	if upcomingBus.NHits >= 2 {
		message.Message = fmt.Sprintf("Prochain bus dans %d min, le suivant dans %d min",
			getDelay(&upcomingBus.Records[0].Information.Departure),
			getDelay(&upcomingBus.Records[1].Information.Departure))
	} else if upcomingBus.NHits == 1 {
		message.Message = fmt.Sprintf("Prochain bus dans %d", getDelay(&upcomingBus.Records[0].Information.Departure))
	}

	json.NewEncoder(w).Encode(message)
}

// Return delay before departure in minutes
func getDelay(departure *time.Time) int {
	return int(departure.Sub(time.Now().UTC()).Minutes())
}
