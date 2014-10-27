package models

import (
	"strconv"
)

type StopTimes struct {
	Records []StopTime
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

func (rs Records) MapToStopTimes() StopTimes {
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
