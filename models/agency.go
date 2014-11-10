package models

type Agencies []Agency

type Agency struct {
	Id int `gorm:"column:id"`
	Key string `gorm:"column:agency_key"`
	AgencyId string `gorm:"column:agency_id"`
	Name string `gorm:"column:agency_name"`
	Url string `gorm:"column:agency_url"`
	Timezone string `gorm:"column:agency_timezone"`
	Lang string `gorm:"column:agency_lang"`
}

type AgencyImportRow struct {
	Key string
	AgencyId string
	Name string
	Url string
	Timezone string
	Lang string
}

type JSONAgencies []JSONAgency

type JSONAgency struct {
	Key string `json:"key"`
	AgencyId string `json:"agencyId"`
	Name string `json:"name"`
	Url string `json:"url"`
	Timezone string `json:"timezone"`
	Lang string `json:"lang"`
}

func (a *Agency) ToJSONAgency() *JSONAgency {
	ja := JSONAgency{
		a.Key,
		a.AgencyId,
		a.Name,
		a.Url,
		a.Timezone,
		a.Lang,
	}

	return &ja
}

func (as *Agencies) ToJSONAgencies() *JSONAgencies {

	jas := make(JSONAgencies, len(*as))

	for i, a := range *as {
		jas[i] = *a.ToJSONAgency()
	}

	return &jas
}

