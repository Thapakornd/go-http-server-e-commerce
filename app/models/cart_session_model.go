package models

import "gorm.io/gorm"

type CartSession struct {
	gorm.Model
	CreatedBy uint64      `json:"user_id" validate:"required"`
	User      User        `gorm:"foreignKey:CreatedBy"`
	CartItems *[]CartItem `gorm:"foreignKey:CartSessionID"`
	Status    string      `json:"status" validate:"required"`
}

type APIAddItem struct {
	UserID    uint64 `json:"user_id" validate:"required"`
	ProductID uint64 `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
