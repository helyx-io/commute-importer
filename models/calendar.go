package models

import(
	"time"
)

type Calendars []Calendar

type Calendar struct {
	AgencyKey string `gorm:"column:agency_key"`
	ServiceId int `gorm:"column:service_id"`
	Monday bool `gorm:"column:monday"`
	Tuesday bool `gorm:"column:tuesday"`
	Wednesday bool `gorm:"column:wednesday"`
	Thursday bool `gorm:"column:thursday"`
	Friday bool `gorm:"column:friday"`
	Saturday bool `gorm:"column:saturday"`
	Sunday bool `gorm:"column:sunday"`
	StartDate time.Time `gorm:"column:start_date"`
	EndDate time.Time `gorm:"column:end_date"`
}

type JSONCalendars []JSONCalendar

type JSONCalendar struct {
	AgencyKey string `json:"agencyKey"`
	ServiceId int `json:"serviceId"`
	Monday bool `json:"monday"`
	Tuesday bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday bool `json:"thursday"`
	Friday bool `json:"friday"`
	Saturday bool `json:"saturday"`
	Sunday bool `json:"sunday"`
	StartDate JSONDate `json:"startDate"`
	EndDate JSONDate `json:"endDate"`
}

type CalendarImportRow struct {
	AgencyKey string
	ServiceId int
	Monday int
	Tuesday int
	Wednesday int
	Thursday int
	Friday int
	Saturday int
	Sunday int
	StartDate string
	EndDate string
}

func (c *Calendar) ToJSONCalendar() *JSONCalendar {
	jc := JSONCalendar{
		c.AgencyKey,
		c.ServiceId,
		c.Monday,
		c.Tuesday,
		c.Wednesday,
		c.Thursday,
		c.Friday,
		c.Saturday,
		c.Sunday,
		JSONDate(c.StartDate),
		JSONDate(c.EndDate),
	}

	return &jc
}

func (cs *Calendars) ToJSONCalendars() *JSONCalendars {

	jcs := make(JSONCalendars, len(*cs))

	for i, c := range *cs {
		jcs[i] = *c.ToJSONCalendar()
	}

	return &jcs
}

