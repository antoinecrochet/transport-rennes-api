package opendatasoft

const (
	endpointBusLineAndBusStopAndDestination = "/api/records/1.0/search/?dataset=tco-bus-circulation-passages-tr" +
		"&q=&facet=nomcourtligne&facet=destination&facet=precision&facet=arrte&facet=nomarret" +
		"&refine.nomcourtligne=%s&refine.nomarret=%s&refine.destination=%s&refine.precision=Temps+réel"

	endpointBusLineAndBusStop = "/api/records/1.0/search/?dataset=tco-bus-circulation-passages-tr" +
		"&q=&facet=nomcourtligne&facet=precision&facet=arrte&facet=nomarret&refine.nomcourtligne=%s" +
		"&refine.nomarret=%s&refine.precision=Temps+réel"

	endpointBusStop = "/api/records/1.0/search/?dataset=tco-bus-circulation-passages-tr" +
		"&q=&facet=precision&facet=nomarret&refine.nomarret=%s&refine.precision=Temps+réel"
)
