package models

//import(
//	"time"
//)

type StopTimes []StopTime

type StopTime struct {
	TripId int `gorm:"column:trip_id"`
	ArrivalTime string `gorm:"column:arrival_time"`
	DepartureTime string `gorm:"column:departure_time"`
	StopId int `gorm:"column:stop_id"`
	StopSequence int `gorm:"column:stop_sequence"`
	StopHeadSign string `gorm:"column:stop_head_sign"`
	PickupType int `gorm:"column:pickup_time"`
	DropOffType int `gorm:"column:drop_off_type"`
	//	ShapeDistTraveled string `gorm:"column:shape_dist_traveled"`
}

type StopTimeImportRow struct {
	TripId int
	ArrivalTime string
	DepartureTime string
	StopId int
	StopSequence int
	StopHeadSign string
	PickupType int
	DropOffType int
	//	ShapeDistTraveled string
}

type JSONStopTimes []JSONStopTime

type JSONStopTime struct {
	TripId int `json:"tripId"`
	ArrivalTime string `json:"arrivalTime"`
	DepartureTime string `json:"departureTime"`
	StopId int `json:"stopId"`
	StopSequence int `json:"stopSequence"`
	StopHeadSign string `json:"stopHeadSign"`
	PickupType int `json:"pickupType"`
	DropOffType int `json:"dropOffType"`
	//	ShapeDistTraveled string `json:"shapeDistTraveled"`
}


func (st *StopTime) ToJSONStopTime() *JSONStopTime {
	jst := JSONStopTime{
		st.TripId,
		st.ArrivalTime,
		st.DepartureTime,
		st.StopId,
		st.StopSequence,
		st.StopHeadSign,
		st.PickupType,
		st.DropOffType,
		//	st.ShapeDistTraveled,
	}

	return &jst
}

func (sts *StopTimes) ToJSONStopTimes() *JSONStopTimes {

	jsts := make(JSONStopTimes, len(*sts))

	for i, st := range *sts {
		jsts[i] = *st.ToJSONStopTime()
	}

	return &jsts
}

