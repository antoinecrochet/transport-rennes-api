package skill

import (
	"os"
	"sort"
	"time"

	db "github.com/antoinecrochet/transport-rennes-be/internal/dynamoDB"
	"github.com/antoinecrochet/transport-rennes-be/internal/opendatasoft"
	"github.com/dasjott/alexa-sdk-go"
)

var odsClient *opendatasoft.OpendatasoftClient

func Initialize() {
	config := opendatasoft.ODSConfig{BaseUrl: os.Getenv("OPENDATASOFT_BASE_URL")}
	odsClient = opendatasoft.New(config)
}

var Handlers = alexa.IntentHandlers{
	"LaunchRequest":    welcome,
	"HelloWorldIntent": welcome,
	"AMAZON.HelpIntent": func(c *alexa.Context) {
		c.Ask(c.T("HELP_MSG"))
	},
	"AMAZON.FallbackIntent": func(c *alexa.Context) {
		c.Ask(c.T("ERROR_MSG"))
	},
	"UpcomingBusIntent":         upcomingBus,
	"UpcomingFavoriteBusIntent": upcomingFavoriteBus,
	"AddToFavoriteIntent":       addFavorite,
	"DeleteFavoriteIntent":      deleteFavorite,
	"AMAZON.StopIntent":         bye,
	"AMAZON.CancelIntent":       bye,
	"SessionEndedRequest":       bye,
}

func welcome(c *alexa.Context) {
	c.Ask(c.T("WELCOME_MSG"))
}

func bye(c *alexa.Context) {
	c.Tell(c.T("GOODBYE_MSG"))
}

func upcomingBus(c *alexa.Context) {
	busStop := c.Slot("busstop")
	busName := c.Slot("bus")

	upcomingBusCommon(c, busStop.Value, busName.Value)
}

func upcomingFavoriteBus(c *alexa.Context) {
	busName := c.Slot("bus")

	favoriteBusStop, err := db.GetFavoriteBusStop(c.System.User.ID)
	if err != nil {
		c.Tell(c.T("FAVORITE_UNAVAILABLE"))
		return
	}

	if favoriteBusStop == nil {
		c.Tell(c.T("NO_FAVORITE"))
		return
	}

	upcomingBusCommon(c, favoriteBusStop.BusStop, busName.Value)
}

func upcomingBusCommon(c *alexa.Context, busStop string, busName string) {
	upcomingBus := odsClient.GetUpcomingBus(busName, busStop, "")

	// If no bus left
	if upcomingBus.NHits == 0 {
		c.Tell(c.TR("NO_BUS_AVAILABLE", alexa.R{"bus": busName, "busstop": busStop}))
		return
	}

	// Too many available buses (alexa asks for more precised request)
	if upcomingBus.NHits >= 10 {
		c.Ask(c.TR("TOO_MANY_BUSES", alexa.R{"busstop": busStop}))
		return
	}

	// sort records by departure time
	sort.SliceStable(upcomingBus.Records, func(i, j int) bool {
		return upcomingBus.Records[i].Information.Departure.Before(upcomingBus.Records[j].Information.Departure)
	})

	// store records by destionation
	x := make(map[string][]opendatasoft.UpcomingBusRecord)
	for _, record := range upcomingBus.Records {
		x[record.Information.Destination] = append(x[record.Information.Destination], record)
	}

	// Generate message
	var message string
	// Loop over destination
	for destination, busList := range x {
		if len(busList) >= 2 {
			message = message + c.TR("UPCOMING_TWO_BUSES_MSG",
				alexa.R{"bus": busList[0].Information.BusLineName,
					"busstop":     busList[0].Information.StopName,
					"destination": destination,
					"dep1":        getDelay(&busList[0].Information.Departure),
					"dep2":        getDelay(&busList[1].Information.Departure)})
		} else if len(busList) == 1 {
			message = message + c.TR("UPCOMING_ONE_BUS_MSG",
				alexa.R{"bus": busList[0].Information.BusLineName,
					"busstop":     busList[0].Information.StopName,
					"destination": destination,
					"dep1":        getDelay(&busList[0].Information.Departure)})
		}
	}

	// Send the final message
	c.Tell(message)
}

// Return delay before departure in minutes
func getDelay(departure *time.Time) int {
	return int(departure.Sub(time.Now().UTC()).Minutes())
}

func addFavorite(c *alexa.Context) {
	busStop := c.Slot("busstop")

	err := db.AddFavoriteBusStop(c.System.User.ID, busStop.Value)
	if err != nil {
		c.Tell(c.T("FAVORITE_UNAVAILABLE"))
		return
	}

	// Send the final message
	c.Tell(c.TR("FAVORITE_SAVED", alexa.R{"busstop": busStop.Value}))
}

func deleteFavorite(c *alexa.Context) {
	favoriteBusStop, err := db.GetFavoriteBusStop(c.System.User.ID)
	if err != nil {
		c.Tell(c.T("FAVORITE_UNAVAILABLE"))
		return
	}

	if favoriteBusStop == nil {
		c.Tell(c.T("NO_FAVORITE"))
		return
	}

	err = db.DeleteFavoriteBusStop(c.System.User.ID)
	if err != nil {
		c.Tell(c.T("FAVORITE_UNAVAILABLE"))
		return
	}

	// Send the final message
	c.Tell(c.T("FAVORITE_DELETED"))
}
