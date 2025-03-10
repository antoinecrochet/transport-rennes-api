package service

import (
	"fmt"
	"sort"

	"github.com/antoinecrochet/transport-rennes-api/internal/core/model"
	"github.com/antoinecrochet/transport-rennes-api/internal/core/port"
)

type Search struct {
	dataProvider port.DataProviderPort
}

func NewSearch(dataProvider port.DataProviderPort) *Search {
	return &Search{
		dataProvider: dataProvider,
	}
}

func (s *Search) SearchUpcomingBus(busLine string, stop string, destination string) (*model.SearchResult, error) {
	if stop == "" {
		return nil, fmt.Errorf("stop bus is mandatory")
	}

	searchResult, err := s.dataProvider.SearchUpcomingPublicTransports(model.Search{BusLine: busLine, BusStop: stop, Destination: destination})
	// sort records by departure time
	sort.SliceStable(searchResult.Hits, func(i, j int) bool {
		return searchResult.Hits[i].Departure.Before(searchResult.Hits[j].Departure)
	})

	return &searchResult, err
}
