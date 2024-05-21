package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartSession   CartSession `gorm:"foreignKey:CartSessionID"`
	Product       Product     `gorm:"foreignKey:ProductID"`
	CartSessionID uint64
	ProductID     uint64
	Quantity      int
}
