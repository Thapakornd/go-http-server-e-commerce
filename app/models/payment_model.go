package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID      uint64
	User        User `gorm:"foreignKey:UserID"`
	PaymentType string
	AccountNo   string
	Expiry      time.Time
}
