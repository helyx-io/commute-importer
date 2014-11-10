package models

type Transfers []Transfer

type Transfer struct {
	Id int `gorm:"column:id"`
	AgencyKey string `gorm:"column:agency_key"`
	FromStopId string `gorm:"column:from_stop_id"`
	ToStopId string `gorm:"column:to_stop_id"`
	TransferType int `gorm:"column:transfer_type"`
	MinTransferTime int `gorm:"column:min_transfer_time"`
}

type TransferImportRow struct {
	AgencyKey string
	FromStopId string
	ToStopId string
	TransferType int
	MinTransferTime int
}

type JSONTransfers []JSONTransfer

type JSONTransfer struct {
	AgencyKey string `json:"agencyKey"`
	FromStopId string `json:"fromStopId"`
	ToStopId string `json:"toStopId"`
	TransferType int `json:"transferType"`
	MinTransferTime int `json:"minTransferTime"`
}

func (t *Transfer) ToJSONTransfer() *JSONTransfer {
	jt := JSONTransfer{
		t.AgencyKey,
		t.FromStopId,
		t.ToStopId,
		t.TransferType,
		t.MinTransferTime,
	}

	return &jt
}

func (ts *Transfers) ToJSONTransfers() *JSONTransfers {

	jts := make(JSONTransfers, len(*ts))

	for i, t := range *ts {
		jts[i] = *t.ToJSONTransfer()
	}

	return &jts
}

