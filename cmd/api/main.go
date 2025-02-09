package main

import "github.com/antoinecrochet/transport-rennes-api/internal/adapter/primary/api"

func main() {
	app := api.New("config.json")
	app.Start()
}
