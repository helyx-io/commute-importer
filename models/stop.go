package models

type Stops []Stop

type Stop struct {
	StopId int `gorm:"stop_id"`
	StopName string `gorm:"stop_name"`
	StopDesc string `gorm:"stop_desc"`
	StopLat int `gorm:"stop_lat"`
	StopLon int `gorm:"stop_lon"`
	ZoneId string `gorm:"zone_id"`
	StopUrl string `gorm:"stop_url"`
	LocationType int `gorm:"location_type"`
	ParentStation int `gorm:"parent_station"`
	//	StopTimzone string `gorm:"stop_timezone"`
}

type StopImportRow struct {
	StopId int
	StopCode string
	StopName string
	StopDesc string
	StopLat float64
	StopLon float64
	ZoneId string
	StopUrl string
	LocationType int
	ParentStation int
	//	StopTimzone string
}

type JSONStops []JSONStop

type JSONStop struct {
	StopId int `json:"stopId"`
	StopName string `json:"stopName"`
	StopDesc string `json:"stopDesc"`
	StopLat int `json:"stopLat"`
	StopLon int `json:"stopLon"`
	ZoneId string `json:"zoneId"`
	StopUrl string `json:"stopUrl"`
	LocationType int `json:"locationType"`
	ParentStation int `json:"parentStation"`
	//	StopTimzone string `json:"stopTimezone"`
}

func (s *Stop) ToJSONStop() *JSONStop {
	js := JSONStop{
		s.StopId,
		s.StopName,
		s.StopDesc,
		s.StopLat,
		s.StopLon,
		s.ZoneId,
		s.StopUrl,
		s.LocationType,
		s.ParentStation,
		//	s.StopTimzone,
	}

	return &js
}

func (ss *Stops) ToJSONStops() *JSONStops {

	jss := make(JSONStops, len(*ss))

	for i, s := range *ss {
		jss[i] = *s.ToJSONStop()
	}

	return &jss
}

