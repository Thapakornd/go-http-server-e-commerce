package models

import "gorm.io/gorm"

type OrderStatus struct {
	gorm.Model
	Order       *[]Order `gorm:"foreignKey:ID"`
	DisplayText string   `json:"display_text" validate:"required"`
}
