package data_provider

import (
	"github.com/antoinecrochet/transport-rennes-api/internal/core/model"
	"github.com/antoinecrochet/transport-rennes-api/opendatasoft"
)

type Opendatasoft struct {
	client opendatasoft.OpendatasoftClientInterface
}

func New(configurationFile string) *Opendatasoft {
	return &Opendatasoft{
		client: opendatasoft.New(configurationFile),
	}
}

func (ods *Opendatasoft) SearchUpcomingPublicTransports(search model.Search) (model.SearchResult, error) {
	result, err := ods.client.SearchUpcomingBus(search.BusStop, search.BusLine, search.Destination)
	if err != nil {
		return model.SearchResult{}, err
	}

	return model.SearchResult{
		Count: result.NHits,
		Hits:  convert(result.Records),
	}, nil
}

// Convert record from opendatasoft into domain model
func convert(odsRecords []opendatasoft.UpcomingBusRecord) []model.PublicTransport {
	result := make([]model.PublicTransport, len(odsRecords))
	for i, record := range odsRecords {
		result[i] = model.PublicTransport{
			BusLineName: record.Information.BusLineName,
			BusStopName: record.Information.StopName,
			Departure:   record.Information.Departure,
			Destination: record.Information.Destination,
		}
	}
	return result
}
