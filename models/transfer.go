package models

type Transfers []Transfer

type Transfer struct {
	AgencyKey string `bson:"agency_key" json:"agencyKey" gorm:"column:agency_key"`
	FromStopId string `bson:"from_stop_id" json:"fromStopId" gorm:"column:from_stop_id"`
	ToStopId string `bson:"to_stop_id" json:"toStopId" gorm:"column:to_stop_id"`
	TransferType int `bson:"transfer_type" json:"transferType" gorm:"column:transfer_type"`
	MinTransferTime int `bson:"min_transfer_time" json:"minTransferTime" gorm:"column:min_transfer_time"`
}
