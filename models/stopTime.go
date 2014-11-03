package models

type StopTimes []StopTime

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
