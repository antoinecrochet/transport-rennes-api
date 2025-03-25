package data_provider

import (
	"errors"
	"testing"

	"time"

	"github.com/antoinecrochet/transport-rennes-api/internal/core/model"
	"github.com/antoinecrochet/transport-rennes-api/opendatasoft"
	mock "github.com/antoinecrochet/transport-rennes-api/opendatasoft/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSearchUpcomingPublicTransports_ReturnsData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockOpendatasoftClientInterface(ctrl)
	ods := &Opendatasoft{client: mockClient}

	search := model.Search{
		BusStop:     "Stop1",
		BusLine:     "Line1",
		Destination: "Destination1",
	}

	expectedResult := &opendatasoft.UpcomingBus{
		NHits: 1,
		Records: []opendatasoft.UpcomingBusRecord{
			{
				Information: opendatasoft.UpcomingBusInformation{
					BusLineName: "Line1",
					StopName:    "Stop1",
					Departure:   time.Now(),
					Destination: "Destination1",
				},
			},
		},
	}

	mockClient.EXPECT().SearchUpcomingBus(search.BusStop, search.BusLine, search.Destination).Return(expectedResult, nil)

	result, err := ods.SearchUpcomingPublicTransports(search)
	assert.NoError(t, err)
	assert.Equal(t, 1, result.Count)
	assert.Equal(t, "Line1", result.Hits[0].BusLineName)
	assert.Equal(t, "Stop1", result.Hits[0].BusStopName)
	assert.NotNil(t, result.Hits[0].Departure)
	assert.Equal(t, "Destination1", result.Hits[0].Destination)
}

func TestSearchUpcomingPublicTransports_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockOpendatasoftClientInterface(ctrl)
	ods := &Opendatasoft{client: mockClient}

	search := model.Search{
		BusStop:     "Stop1",
		BusLine:     "Line1",
		Destination: "Destination1",
	}

	mockClient.EXPECT().SearchUpcomingBus(search.BusStop, search.BusLine, search.Destination).Return(&opendatasoft.UpcomingBus{}, errors.New("service unavailable"))

	_, err := ods.SearchUpcomingPublicTransports(search)
	assert.Error(t, err)
	assert.Equal(t, "service unavailable", err.Error())
}
