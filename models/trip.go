package models

type Trips []Trip

type Trip struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey" gorm:"column:agency_key"`
	RouteId string `bson:"route_id" json:"routeId" gorm:"column:route_id"`
	ServiceId int `bson:"service_id" json:"serviceId" gorm:"column:service_id"`
	TripId string `bson:"trip_id" json:"tripId" gorm:"column:trip_id"`
	TripHeadsign string `bson:"trip_headsign" json:"tripHeadsign" gorm:"column:trip_headsign"`
	DirectionId int `bson:"direction_id" json:"directionId" gorm:"column:direction_id"`
	BlockId string `bson:"block_id" json:"blockId" gorm:"column:block_id"`
	ShapeId string `bson:"shape_id" json:"shapeId" gorm:"column:shape_id"`
}
