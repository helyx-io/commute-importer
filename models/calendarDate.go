package models

type CalendarDates []CalendarDate

type CalendarDate struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey" gorm:"column:agency_key"`
	ServiceId int `bson:"service_id" json:"serviceId" gorm:"column:service_id"`
	Date string `bson:"date" json:"date" gorm:"column:date"`
	ExceptionType int `bson:"exception_type" json:"exceptionType" gorm:"column:exception_type"`
}
