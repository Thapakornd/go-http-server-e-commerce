package models

import (
	"gorm.io/gorm"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Discount    float64
	CategoryID  uint64
	Category    Category `gorm:"foreignKey:CategoryID"`
	CartItem    *CartItem
	gorm.Model
}
