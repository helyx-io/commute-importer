package models

import(
	"time"
)

type CalendarDates []CalendarDate

type CalendarDate struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey" gorm:"column:agency_key"`
	ServiceId int `bson:"service_id" json:"serviceId" gorm:"column:service_id"`
	Date time.Time `bson:"date" json:"date" gorm:"column:date"`
	ExceptionType int `bson:"exception_type" json:"exceptionType" gorm:"column:exception_type"`
}

type CalendarDateImportRow struct {
	AgencyKey string `gorm:"column:agency_key"`
	ServiceId int `gorm:"column:service_id"`
	Date string `gorm:"column:date"`
	ExceptionType int `gorm:"column:exception_type"`
}
