package service

import (
	"github.com/jinzhu/gorm"
	"database/sql"
)

type StopTimeService interface {
	RemoveByAgencyKey(agencyKey string) (error)
}
