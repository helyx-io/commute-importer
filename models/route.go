package models

type Routes []Route

type Route struct {
	Id int `gorm:"column:id"`
	AgencyKey string `gorm:"column:agency_key"`
	RouteId string `gorm:"column:route_id"`
	AgencyId string `gorm:"column:agency_id"`
	RouteShortName string `gorm:"column:route_short_name"`
	RouteLongName string `gorm:"column:route_long_name"`
	RouteDesc string `gorm:"column:route_desc"`
	RouteType int `gorm:"column:route_type"`
	RouteUrl string `gorm:"column:route_url"`
	RouteColor string `gorm:"column:route_color"`
	RouteTextColor string `gorm:"column:route_text_color"`
}

type RouteImportRow struct {
	AgencyKey string
	RouteId string
	AgencyId string
	RouteShortName string
	RouteLongName string
	RouteDesc string
	RouteType int
	RouteUrl string
	RouteColor string
	RouteTextColor string
}

type JSONRoutes []JSONRoute

type JSONRoute struct {
	AgencyKey string `json:"agencyKey"`
	RouteId string `json:"routeId"`
	AgencyId string `json:"agencyId"`
	RouteShortName string `json:"routeShortName"`
	RouteLongName string `json:"routeLongName"`
	RouteDesc string `json:"routeDesc"`
	RouteType int `json:"routeType"`
	RouteUrl string `json:"routeUrl"`
	RouteColor string `json:"routeColor"`
	RouteTextColor string `json:"routeTextColor"`
}

func (r *Route) ToJSONRoute() *JSONRoute {
	jr := JSONRoute{
		r.AgencyKey,
		r.RouteId,
		r.AgencyId,
		r.RouteShortName,
		r.RouteLongName,
		r.RouteDesc,
		r.RouteType,
		r.RouteUrl,
		r.RouteColor,
		r.RouteTextColor,
	}

	return &jr
}

func (rs *Routes) ToJSONRoutes() *JSONRoutes {

	jrs := make(JSONRoutes, len(*rs))

	for i, r := range *rs {
		jrs[i] = *r.ToJSONRoute()
	}

	return &jrs
}

