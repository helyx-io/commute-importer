package models

import (
	"strconv"
)

type Stops struct {
	Records []Stop
}

type Stop struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey"`
	StopId string `bson:"stop_id" json:"stopId"`
	StopCode string `bson:"stop_code" json:"stopCode"`
	StopName string `bson:"stop_name" json:"stopName"`
	StopDesc string `bson:"stop_desc" json:"stopDesc"`
	StopLat int `bson:"stop_lat" json:"stopLat"`
	StopLon int `bson:"stop_lon" json:"stopLon"`
	ZoneId string `bson:"zone_id" json:"zoneId"`
	StopUrl string `bson:"stop_url" json:"stopUrl"`
	LocationType int `bson:"location_type" json:"locationType"`
	ParentStation string `bson:"parent_station" json:"paretnStation"`
//	StopTimzone string `bson:"stop_timezone" json:"stopTimezone"`
}

func (rs *Records) MapToStops() Stops {
	var st = Stops{ make([]Stop, len(rs.Records)) }

	for i, record := range rs.Records {
		stopLat, _ := strconv.Atoi(record[4])
		stopLon, _ := strconv.Atoi(record[5])
		locationType, _ := strconv.Atoi(record[8])
		st.Records[i] = Stop{
			"RATP",
			record[0],
			record[1],
			record[2],
			record[3],
			stopLat,
			stopLon,
			record[6],
			record[7],
			locationType,
			record[9],
		}
	}

	return st
}
