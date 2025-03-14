package port

import "github.com/antoinecrochet/transport-rennes-api/internal/core/model"

//go:generate mockgen -source=secondary.go -destination=mock/secondary.go

type DataProviderPort interface {
	SearchUpcomingPublicTransports(search model.Search) (model.SearchResult, error)
}
