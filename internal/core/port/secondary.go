package port

import "github.com/antoinecrochet/transport-rennes-api/internal/core/model"

type DataProviderPort interface {
	SearchUpcomingPublicTransports(search model.Search) (model.SearchResult, error)
}
