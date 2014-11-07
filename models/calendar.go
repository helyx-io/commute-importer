package models

import(
	"time"
)

type Calendars []Calendar

type Calendar struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey" gorm:"column:agency_key"`
	ServiceId int `bson:"service_id" json:"serviceId" gorm:"column:service_id"`
	Monday int `bson:"monday" json:"monday" gorm:"column:monday"`
	Tuesday int `bson:"tuesday" json:"tuesday" gorm:"column:tuesday"`
	Wednesday int `bson:"wednesday" json:"wednesday" gorm:"column:wednesday"`
	Thursday int `bson:"thursday" json:"thursday" gorm:"column:thursday"`
	Friday int `bson:"friday" json:"friday" gorm:"column:friday"`
	Saturday int `bson:"saturday" json:"saturday" gorm:"column:saturday"`
	Sunday int `bson:"sunday" json:"sunday" gorm:"column:sunday"`
	StartDate time.Time `bson:"start_date" json:"startDate" gorm:"column:start_date"`
	EndDate time.Time `bson:"end_date" json:"endDate" gorm:"column:end_date"`
}

type CalendarImportRow struct {
	AgencyKey string `gorm:"column:agency_key"`
	ServiceId int `gorm:"column:service_id"`
	Monday int `gorm:"column:monday"`
	Tuesday int `gorm:"column:tuesday"`
	Wednesday int `gorm:"column:wednesday"`
	Thursday int `gorm:"column:thursday"`
	Friday int `gorm:"column:friday"`
	Saturday int `gorm:"column:saturday"`
	Sunday int `gorm:"column:sunday"`
	StartDate string `gorm:"column:start_date"`
	EndDate string `gorm:"column:end_date"`
}
