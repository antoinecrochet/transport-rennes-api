package main

import (
	"github.com/antoinecrochet/transport-rennes-api/internal/adapter/primary/api"
	"github.com/antoinecrochet/transport-rennes-api/internal/adapter/secondary/data_provider"
	"github.com/antoinecrochet/transport-rennes-api/internal/core/service"
)

func main() {
	// init data provider
	dataProvider := data_provider.NewOpendatasoftClient("config.json")

	// init search service
	search := service.NewSearch(dataProvider)

	// init application and start the server
	app := api.NewApplication(search)
	app.Start()
}
