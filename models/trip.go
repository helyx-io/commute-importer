package models

type Trips []Trip

type Trip struct {
	RouteId int `bson:"route_id" json:"routeId" gorm:"column:route_id"`
	ServiceId int `bson:"service_id" json:"serviceId" gorm:"column:service_id"`
	TripId int `bson:"trip_id" json:"tripId" gorm:"column:trip_id"`
	TripHeadsign string `bson:"trip_headsign" json:"tripHeadsign" gorm:"column:trip_headsign"`
	DirectionId int `bson:"direction_id" json:"directionId" gorm:"column:direction_id"`
	BlockId string `bson:"block_id" json:"blockId" gorm:"column:block_id"`
	ShapeId string `bson:"shape_id" json:"shapeId" gorm:"column:shape_id"`
}

type TripImportRow struct {
	RouteId int
	ServiceId int
	TripId int
	TripHeadsign string
	DirectionId int
	BlockId string
	ShapeId string
}

type JSONTrips []JSONTrip

type JSONTrip struct {
	RouteId int `json:"routeId"`
	ServiceId int `json:"serviceId"`
	TripId int `json:"tripId"`
	TripHeadsign string `json:"tripHeadsign"`
	DirectionId int `json:"directionId"`
	BlockId string `json:"blockId"`
	ShapeId string `json:"shapeId"`
}


func (t *Trip) ToJSONTrip() *JSONTrip {
	jt := JSONTrip{
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

