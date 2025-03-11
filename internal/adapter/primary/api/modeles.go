package api

import (
	"time"
)

type SearchUpcomingBusResponse struct {
	Message string        `json:"message"`
	Count   int           `json:"totalCount"`
	Hits    []HitResponse `json:"hits"`
}

type HitResponse struct {
	BusLineName string    `json:"busline"`
	BusStopName string    `json:"stop"`
	Departure   time.Time `json:"departure"`
	Destination string    `json:"destination"`
}

type SearchBus struct {
	BusLine     string `json:"busline"`
	Stop        string `json:"stop"`
	Destination string `json:"destination"`
}

type Error struct {
	ErrorMessage string `json:"error"`
}
