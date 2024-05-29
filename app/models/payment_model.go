package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID      uint64 `json:"user_id" validate:"required"`
	User        User   `gorm:"foreignKey:UserID"`
	PaymentType string `json:"type" validate:"required"`
	AccountNo   string `json:"account_no" validate:"required"`
	Expiry      string `json:"expiry" validate:"required"`
}
