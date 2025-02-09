package opendatasoft

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const endpointAll string = "/api/records/1.0/search/?dataset=tco-bus-circulation-passages-tr" +
	"&q=&facet=nomcourtligne&facet=destination&facet=precision&facet=arrte&facet=nomarret" +
	"&refine.nomcourtligne=%s&refine.nomarret=%s&refine.destination=%s&refine.precision=Temps+réel"

const endpointAllButDestionation string = "/api/records/1.0/search/?dataset=tco-bus-circulation-passages-tr" +
	"&q=&facet=nomcourtligne&facet=precision&facet=arrte&facet=nomarret&refine.nomcourtligne=%s" +
	"&refine.nomarret=%s&refine.precision=Temps+réel"

const endpointOnlyStopName string = "/api/records/1.0/search/?dataset=tco-bus-circulation-passages-tr" +
	"&q=&facet=precision&facet=nomarret&refine.nomarret=%s&refine.precision=Temps+réel"

type OpendatasoftClient struct {
	client *http.Client
	config ODSConfig
}

// Create new instance of OpendatasoftClient
func New(config ODSConfig) *OpendatasoftClient {
	return &OpendatasoftClient{
		client: &http.Client{Timeout: 10 * time.Second},
		config: config,
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
		query = fmt.Sprintf(endpointOnlyStopName, url.QueryEscape(stopName))
	} else if destination == "" && busLineName != "" {
		// Seach bus by bus line name and stop name
		query = fmt.Sprintf(endpointAllButDestionation, url.QueryEscape(busLineName), url.QueryEscape(stopName))
	} else {
		// Seach bus by bus line name, destination and stop name
		query = fmt.Sprintf(endpointAll, url.QueryEscape(busLineName), url.QueryEscape(stopName), url.QueryEscape(destination))
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
