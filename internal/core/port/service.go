package port

import "github.com/antoinecrochet/transport-rennes-api/internal/core/model"

type SearchBus interface {
	SearchUpcomingBus(busLine string, stop string, destination string) (model.SearchResult, error)
}
