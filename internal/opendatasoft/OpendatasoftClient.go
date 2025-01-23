package opendatasoft

import (
	"encoding/json"
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
	httpClient := OpendatasoftClient{}
	httpClient.client = &http.Client{Timeout: 10 * time.Second}
	httpClient.config = config

	return &httpClient
}

func (ods *OpendatasoftClient) GetUpcomingBus(busLineName string, stopName string, destination string) *UpcomingBus {
	var upcomingBus UpcomingBus

	if destination == "" && busLineName == "" {
		// Search bus by stop name
		ods.getRequest(fmt.Sprintf(endpointOnlyStopName, url.QueryEscape(stopName)), &upcomingBus)
	} else if destination == "" && busLineName != "" {
		// Seach bus by bus line name and stop name
		ods.getRequest(fmt.Sprintf(endpointAllButDestionation, url.QueryEscape(busLineName), url.QueryEscape(stopName)), &upcomingBus)
	} else {
		// Seach bus by bus line name, destination and stop name
		ods.getRequest(fmt.Sprintf(endpointAll, url.QueryEscape(busLineName), url.QueryEscape(stopName), url.QueryEscape(destination)), &upcomingBus)
	}

	return &upcomingBus
}

// Get request to opendatasoft api
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
