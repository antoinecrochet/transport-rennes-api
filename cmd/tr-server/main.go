package main

import (
	"log"
	"net/http"

	api "github.com/antoinecrochet/transport-rennes-api/internal/transport-rennes-api"
)

func main() {
	api.InitializeServer()

	router := api.InitializeRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
