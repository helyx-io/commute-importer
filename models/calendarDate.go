package models

import(
	"time"
)

type CalendarDates []CalendarDate

type CalendarDate struct {
	ServiceId int `gorm:"column:service_id"`
	Date time.Time `gorm:"column:date"`
	ExceptionType int `gorm:"column:exception_type"`
}

type CalendarDateImportRow struct {
	ServiceId int
	Date string
	ExceptionType int
}

type JSONCalendarDates []JSONCalendarDate

type JSONCalendarDate struct {
	ServiceId int `json:"serviceId"`
	Date JSONDate `json:"date"`
	ExceptionType int `json:"exceptionType"`
}

func (c *CalendarDate) ToJSONCalendarDate() *JSONCalendarDate {
	jc := JSONCalendarDate{
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

