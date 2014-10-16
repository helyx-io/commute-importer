package models

import (
	"strconv"
	"gopkg.in/mgo.v2/bson"
)

type Agency struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	AgencyId string `bson:"agency_id" json:"agencyId"`
	Name string `bson:"agency_name" json:"name"`
	Url string `bson:"agency_url" json:"url"`
	Timezone string `bson:"agency_timezone" json:"timezone"`
	Lang string `bson:"agency_lang" json:"lang"`
	Key string `bson:"agency_key" json:"key"`
}

type Stop struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey"`
	StopId string `bson:"stop_id" json:"stopId"`
	StopCode string `bson:"stop_code" json:"stopCode"`
	StopName string `bson:"stop_name" json:"stopName"`
	StopDesc string `bson:"stop_desc" json:"stopDesc"`
	StopLat string `bson:"stop_lat" json:"stopLat"`
	StopLon string `bson:"stop_lon" json:"stopLon"`
	Loc []string `bson:"lov" json:"loc"`
	ZoneId string `bson:"zone_id" json:"zoneId"`
	StopUrl string `bson:"stop_url" json:"stopUrl"`
	locationType string `bson:"location_type" json:"locationType"`
	ParentStation string `bson:"parent_station" json:"paretnStation"`
	StopTimzone string `bson:"stop_timezone" json:"stopTimezone"`
}

type StopTime struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey"`
	TripId string `bson:"trip_id" json:"tripId"`
	ArrivalTime string `bson:"stop_code" json:"arrivalTime"`
	DepartureTime string `bson:"departure_time" json:"departureTime"`
	StopId string `bson:"stop_id" json:"stopId"`
	StopSequence int `bson:"stop_sequence" json:"stopSequence"`
	StopHeadSign string `bson:"stop_head_sign" json:"stopHeadSign"`
	PickupType int `bson:"pickup_type" json:"pickupType"`
	DropOffType int `bson:"drop_off_type" json:"dropOffType"`
//	ShapeDistTraveled string `bson:"shape_dist_traveled" json:"shapeDistTraveled"`
}

type Records struct {
	Records [][]string
}

type StopTimes struct {
	Records []StopTime
}

type RecordsInserter interface {
	InsertStopTimes(sts *StopTimes) (err error)
}

func (rs *Records) MapToStopTimes() StopTimes {
	var st = StopTimes{ make([]StopTime, len(rs.Records)) }

	for i, record := range rs.Records {
		stopSequence, _ := strconv.Atoi(record[4])
		pickup_type, _ := strconv.Atoi(record[6])
		drop_off_type, _ := strconv.Atoi(record[7])
		st.Records[i] = StopTime{
			"RATP",
			record[0],
			record[1],
			record[2],
			record[3],
			stopSequence,
			record[5],
			pickup_type,
			drop_off_type,
		}
	}

	return st
}
