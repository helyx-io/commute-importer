package models

type Stops []Stop

type Stop struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey"`
	StopId string `bson:"stop_id" json:"stopId"`
	StopName string `bson:"stop_name" json:"stopName"`
	StopDesc string `bson:"stop_desc" json:"stopDesc"`
	StopLat int `bson:"stop_lat" json:"stopLat"`
	StopLon int `bson:"stop_lon" json:"stopLon"`
	ZoneId string `bson:"zone_id" json:"zoneId"`
	StopUrl string `bson:"stop_url" json:"stopUrl"`
	LocationType int `bson:"location_type" json:"locationType"`
	ParentStation int `bson:"parent_station" json:"paretnStation"`
//	StopTimzone string `bson:"stop_timezone" json:"stopTimezone"`
}
