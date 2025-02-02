package opendatasoft

import (
	"encoding/json"
	"log"
	"os"
)

// Json structure
type ODSConfig struct {
	BaseUrl string `json:"base_url"`
}

var config *ODSConfig

// Read config file and store data
func ReadConfigFile(url string) *ODSConfig {
	file, err := os.Open(url)
	decoder := json.NewDecoder(file)

	if err != nil {
		log.Fatal(err)
	}

	err2 := decoder.Decode(&config)
	if err2 != nil {
		log.Fatal("Error while parsing config.json: ", err)
	}

	return config
}
