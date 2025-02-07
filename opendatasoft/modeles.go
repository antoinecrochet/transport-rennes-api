package opendatasoft

import "time"

type UpcomingBus struct {
	NHits   int                 `json:"nhits"`
	Records []UpcomingBusRecord `json:"records"`
}

type UpcomingBusRecord struct {
	DatasetId   string                 `json:"datasetid"`
	RecordId    string                 `json:"recordid"`
	Information UpcomingBusInformation `json:"fields"`
	Timestamp   time.Time              `json:"record_timestamp"`
}

type UpcomingBusInformation struct {
	BusLineName          string    `json:"nomcourtligne"`
	Departure            time.Time `json:"depart"`
	TheoreticalDeparture time.Time `json:"departtheorique"`
	TheoreticalArrival   time.Time `json:"arriveetheorique"`
	Destination          string    `json:"destination"`
	StopName             string    `json:"nomarret"`
	Precision            string    `json:"precision"`
	CourseId             string    `json:"idcourse"`
	ArrivalTime          time.Time `json:"arrivee"`
	Direction            int       `json:"sens"`
	StopId               string    `json:"idarret"`
	BusLineId            string    `json:"idligne"`
}
