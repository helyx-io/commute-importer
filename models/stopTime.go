package models

import(
	"time"
)

type StopTimes []StopTime

type StopTime struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey" gorm:"column:agency_key"`
	TripId string `bson:"trip_id" json:"tripId" gorm:"column:trip_id"`
	ArrivalTime time.Time `bson:"stop_code" json:"arrivalTime" gorm:"column:arrival_time"`
	DepartureTime time.Time `bson:"departure_time" json:"departureTime" gorm:"column:departure_time"`
	StopId string `bson:"stop_id" json:"stopId" gorm:"column:stop_id"`
	StopSequence int `bson:"stop_sequence" json:"stopSequence" gorm:"column:stop_sequence"`
	StopHeadSign string `bson:"stop_head_sign" json:"stopHeadSign" gorm:"column:stop_head_sign"`
	PickupType int `bson:"pickup_type" json:"pickupType" gorm:"column:pickup_time"`
	DropOffType int `bson:"drop_off_type" json:"dropOffType" gorm:"column:drop_off_type"`
	//	ShapeDistTraveled string `bson:"shape_dist_traveled" json:"shapeDistTraveled" gorm:"column:shape_dist_traveled"`
}

type StopTimeImportRow struct {
	AgencyKey string `gorm:"column:agency_key"`
	TripId string `gorm:"column:trip_id"`
	ArrivalTime string `gorm:"column:arrival_time"`
	DepartureTime string `gorm:"column:departure_time"`
	StopId string `gorm:"column:stop_id"`
	StopSequence int `gorm:"column:stop_sequence"`
	StopHeadSign string `gorm:"column:stop_head_sign"`
	PickupType int `gorm:"column:pickup_time"`
	DropOffType int `gorm:"column:drop_off_type"`
	//	ShapeDistTraveled string `gorm:"column:shape_dist_traveled"`
}
