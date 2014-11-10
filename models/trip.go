package models

type Trips []Trip

type Trip struct {
	Id int `gorm:"column:id"`
	AgencyKey string `bson:"agency_key" json:"agencyKey" gorm:"column:agency_key"`
	RouteId string `bson:"route_id" json:"routeId" gorm:"column:route_id"`
	ServiceId int `bson:"service_id" json:"serviceId" gorm:"column:service_id"`
	TripId string `bson:"trip_id" json:"tripId" gorm:"column:trip_id"`
	TripHeadsign string `bson:"trip_headsign" json:"tripHeadsign" gorm:"column:trip_headsign"`
	DirectionId int `bson:"direction_id" json:"directionId" gorm:"column:direction_id"`
	BlockId string `bson:"block_id" json:"blockId" gorm:"column:block_id"`
	ShapeId string `bson:"shape_id" json:"shapeId" gorm:"column:shape_id"`
}

type TripImportRow struct {
	AgencyKey string
	RouteId string
	ServiceId int
	TripId string
	TripHeadsign string
	DirectionId int
	BlockId string
	ShapeId string
}

type JSONTrips []JSONTrip

type JSONTrip struct {
	AgencyKey string `json:"agencyKey"`
	RouteId string `json:"routeId"`
	ServiceId int `json:"serviceId"`
	TripId string `json:"tripId"`
	TripHeadsign string `json:"tripHeadsign"`
	DirectionId int `json:"directionId"`
	BlockId string `json:"blockId"`
	ShapeId string `json:"shapeId"`
}


func (t *Trip) ToJSONTrip() *JSONTrip {
	jt := JSONTrip{
		t.AgencyKey,
		t.RouteId,
		t.ServiceId,
		t.TripId,
		t.TripHeadsign,
		t.DirectionId,
		t.BlockId,
		t.ShapeId,
	}

	return &jt
}

func (ts *Trips) ToJSONTrips() *JSONTrips {

	jts := make(JSONTrips, len(*ts))

	for i, t := range *ts {
		jts[i] = *t.ToJSONTrip()
	}

	return &jts
}

