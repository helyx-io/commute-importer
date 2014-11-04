package models

type Routes []Route

type Route struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey" gorm:"column:agency_key"`
	RouteId string `bson:"route_id" json:"routeId" gorm:"column:route_id"`
	AgencyId string `bson:"agency_id" json:"agencyId" gorm:"column:agency_id"`
	RouteShortName string `bson:"route_short_name" json:"routeShortName" gorm:"column:route_short_name"`
	RouteLongName string `bson:"route_long_name" json:"routeLongName" gorm:"column:route_long_name"`
	RouteDesc string `bson:"route_desc" json:"routeDesc" gorm:"column:route_desc"`
	RouteType int `bson:"route_type" json:"routeType" gorm:"column:route_type"`
	RouteUrl string `bson:"route_url" json:"routeUrl" gorm:"column:route_url"`
	RouteColor string `bson:"route_color" json:"routeColor" gorm:"column:route_color"`
	RouteTextColor string `bson:"route_text_color" json:"routeTextColor" gorm:"column:route_text_color"`
}
