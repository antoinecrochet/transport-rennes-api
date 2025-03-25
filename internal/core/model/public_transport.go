package model

import "time"

type PublicTransport struct {
	BusLineName string
	BusStopName string
	Departure   time.Time
	Destination string
}
