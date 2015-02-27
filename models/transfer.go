package models

type Transfers []Transfer

type Transfer struct {
	Id int `gorm:"column:id"`
	FromStopId int `gorm:"column:from_stop_id"`
	ToStopId int `gorm:"column:to_stop_id"`
	TransferType int `gorm:"column:transfer_type"`
	MinTransferTime int `gorm:"column:min_transfer_time"`
}

type TransferImportRow struct {
	FromStopId int
	ToStopId int
	TransferType int
	MinTransferTime int
}

type JSONTransfers []JSONTransfer

type JSONTransfer struct {
	FromStopId int `json:"fromStopId"`
	ToStopId int `json:"toStopId"`
	TransferType int `json:"transferType"`
	MinTransferTime int `json:"minTransferTime"`
}

func (t *Transfer) ToJSONTransfer() *JSONTransfer {
	jt := JSONTransfer{
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

