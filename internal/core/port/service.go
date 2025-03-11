package port

import "github.com/antoinecrochet/transport-rennes-api/internal/core/model"

type Search interface {
	SearchUpcomingBus(search model.Search) (*model.SearchResult, error)
}
