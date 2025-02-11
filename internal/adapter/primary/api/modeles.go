package api

type Message struct {
	Message string `json:"message"`
}

type SearchBus struct {
	BusLine     string `json:"busline"`
	Stop        string `json:"stop"`
	Destination string `json:"destination"`
}
