package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartSession   CartSession `gorm:"foreignKey:CartSessionID"`
	Product       Product     `gorm:"foreignKey:ProductID"`
	CartSessionID uint64      `json:"cart_id" validate:"required"`
	ProductID     uint64      `json:"product_id" validate:"required"`
	Quantity      int         `json:"quantity" validate:"required"`
}
