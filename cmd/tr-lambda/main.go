package main

import (
	"os"

	skill "github.com/antoinecrochet/transport-rennes-be/internal/alexa-skill"
	db "github.com/antoinecrochet/transport-rennes-be/internal/dynamoDB"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dasjott/alexa-sdk-go"
)

func main() {
	skill.Initialize()
	db.InitializeDynamodbClient()

	alexa.AppID = os.Getenv("ALEXA_APP_ID")
	alexa.Handlers = skill.Handlers
	alexa.LocaleStrings = skill.Locales
	lambda.Start(alexa.Handle)
}
