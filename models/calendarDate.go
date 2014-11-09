package models

import(
	"time"
)

type CalendarDates []CalendarDate

type CalendarDate struct {
	AgencyKey string `gorm:"column:agency_key"`
	ServiceId int `gorm:"column:service_id"`
	Date time.Time `gorm:"column:date"`
	ExceptionType int `gorm:"column:exception_type"`
}

type JSONCalendarDates []JSONCalendarDate

type JSONCalendarDate struct {
	AgencyKey string `json:"agencyKey"`
	ServiceId int `json:"serviceId"`
	Date JSONDate `json:"date"`
	ExceptionType int `json:"exceptionType"`
}

type CalendarDateImportRow struct {
	AgencyKey string
	ServiceId int
	Date string
	ExceptionType int
}

func (c *CalendarDate) ToJSONCalendarDate() *JSONCalendarDate {
	jc := JSONCalendarDate{
		c.AgencyKey,
		c.ServiceId,
		JSONDate(c.Date),
		c.ExceptionType,
	}

	return &jc
}

func (cs *CalendarDates) ToJSONCalendarDates() *JSONCalendarDates {

	jcs := make(JSONCalendarDates, len(*cs))

	for i, c := range *cs {
		jcs[i] = *c.ToJSONCalendarDate()
	}

	return &jcs
}

