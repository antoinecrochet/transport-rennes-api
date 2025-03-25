package service

import (
	"testing"

	"time"

	"github.com/antoinecrochet/transport-rennes-api/internal/core/model"
	mock_port "github.com/antoinecrochet/transport-rennes-api/internal/core/port/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSearchUpcomingBus_ReturnsResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_port.NewMockDataProviderPort(ctrl)
	search := &Search{dataProvider: mockClient}

	request := model.Search{BusLine: "Line1", BusStop: "Stop1", Destination: "Destination1"}
	expectedResult := &model.SearchResult{
		Count: 1,
		Hits: []model.PublicTransport{
			{
				BusLineName: "Line1",
				BusStopName: "Stop1",
				Departure:   time.Now(),
				Destination: "Destination1",
			},
		},
	}

	mockClient.EXPECT().SearchUpcomingPublicTransports(request).Return(*expectedResult, nil)

	result, err := search.SearchUpcomingBus(request)
	assert.NoError(t, err)
	assert.Equal(t, 1, result.Count)
	assert.Equal(t, "Line1", result.Hits[0].BusLineName)
	assert.Equal(t, "Stop1", result.Hits[0].BusStopName)
	assert.NotNil(t, result.Hits[0].Departure)
	assert.Equal(t, "Destination1", result.Hits[0].Destination)
}

func TestSearchUpcomingBus_ReturnsResultSortedByDeparture(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_port.NewMockDataProviderPort(ctrl)
	search := &Search{dataProvider: mockClient}

	request := model.Search{BusLine: "Line1", BusStop: "Stop1", Destination: "Destination1"}
	departure6min := time.Now().Add(time.Minute * 6)
	departure1min := time.Now().Add(time.Minute * 1)
	departure10min := time.Now().Add(time.Minute * 10)
	expectedResult := &model.SearchResult{
		Count: 3,
		Hits: []model.PublicTransport{
			{
				BusLineName: "Line1",
				BusStopName: "Stop1",
				Departure:   departure6min, // departure in 6 minutes
				Destination: "Destination1",
			},
			{
				BusLineName: "Line1",
				BusStopName: "Stop1",
				Departure:   departure1min, // departure in 1 minute
				Destination: "Destination1",
			},
			{
				BusLineName: "Line1",
				BusStopName: "Stop1",
				Departure:   departure10min, // departure in 10 minute
				Destination: "Destination1",
			},
		},
	}

	mockClient.EXPECT().SearchUpcomingPublicTransports(request).Return(*expectedResult, nil)

	result, err := search.SearchUpcomingBus(request)
	assert.NoError(t, err)
	assert.Equal(t, 3, result.Count)
	// Check it has been sorted by departure time
	assert.Equal(t, departure1min, result.Hits[0].Departure)
	assert.Equal(t, departure6min, result.Hits[1].Departure)
	assert.Equal(t, departure10min, result.Hits[2].Departure)
}

func TestSearchUpcomingBusMissingMandatoryParam_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_port.NewMockDataProviderPort(ctrl)
	search := &Search{dataProvider: mockClient}

	request := model.Search{BusLine: "Line1", Destination: "Destination1"}

	_, err := search.SearchUpcomingBus(request)
	assert.Error(t, err)
}
