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

type APIProduct struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Category    string  `json:"category" validate:"required"`
}
