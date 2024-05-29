package models

import "gorm.io/gorm"

type CartSession struct {
	gorm.Model
	CreatedBy uint64
	User      User        `gorm:"foreignKey:CreatedBy"`
	CartItems *[]CartItem `gorm:"foreignKey:CartSessionID"`
	Status    string
}

type APICartSession struct {
	CreatedBy int64
	Status    string
}
