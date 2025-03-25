package opendatasoft

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

//go:generate mockgen -source=service.go -destination=mock/service.go

type OpendatasoftClientInterface interface {
	SearchUpcomingBus(stopName string, busLineName string, destination string) (*UpcomingBus, error)
}

type OpendatasoftClient struct {
	client *http.Client
	config *OpendatasoftConfig
}

// Create new instance of OpendatasoftClient
func New(configurationFile string) *OpendatasoftClient {
	return &OpendatasoftClient{
		client: &http.Client{Timeout: 10 * time.Second},
		config: readConfigurationFile(configurationFile),
	}
}

// Search upcoming bus in real time given bus line, stop name and destination
// StopName is the only required parameter
func (ods *OpendatasoftClient) SearchUpcomingBus(stopName string, busLineName string, destination string) (*UpcomingBus, error) {
	if stopName == "" {
		return nil, errors.New("stopName is required")
	}

	var query string
	if destination == "" && busLineName == "" {
		// Search bus by stop name
		query = fmt.Sprintf(endpointBusStop, url.QueryEscape(stopName))
	} else if destination == "" && busLineName != "" {
		// Seach bus by bus line name and stop name
		query = fmt.Sprintf(endpointBusLineAndBusStop, url.QueryEscape(busLineName), url.QueryEscape(stopName))
	} else {
		// Seach bus by bus line name, destination and stop name
		query = fmt.Sprintf(endpointBusLineAndBusStopAndDestination, url.QueryEscape(busLineName), url.QueryEscape(stopName), url.QueryEscape(destination))
	}

	var upcomingBus UpcomingBus
	err := ods.getRequest(query, &upcomingBus)
	return &upcomingBus, err
}

// HTTP get request to opendatasoft api
func (ods *OpendatasoftClient) getRequest(request string, target interface{}) error {
	req, _ := http.NewRequest("GET", ods.config.BaseUrl+request, nil)
	resp, err := ods.client.Do(req)
	if err != nil {
		log.Panic(err)
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&target)
	return nil
}

// Read configuration file
func readConfigurationFile(path string) *OpendatasoftConfig {
	file, err := os.Open(path)
	decoder := json.NewDecoder(file)

	if err != nil {
		log.Fatal(err)
	}

	var config *OpendatasoftConfig
	err2 := decoder.Decode(&config)
	if err2 != nil {
		log.Fatal("Error while parsing config.json: ", err)
	}

	return config
}
