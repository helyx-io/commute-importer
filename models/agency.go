package models

type Agencies []Agency

type Agency struct {
	Key string `bson:"agency_key" json:"key" gorm:"column:agency_key"`
	Id string `bson:"agency_id" json:"agencyId" gorm:"column:agency_id"`
	Name string `bson:"agency_name" json:"name" gorm:"column:agency_name"`
	Url string `bson:"agency_url" json:"url" gorm:"column:agency_url"`
	Timezone string `bson:"agency_timezone" json:"timezone" gorm:"column:agency_timezone"`
	Lang string `bson:"agency_lang" json:"lang" gorm:"column:agency_lang"`
}

