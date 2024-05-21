package models

import "gorm.io/gorm"

type ShippingMethod struct {
	gorm.Model
	Order         *[]Order `gorn:"foreignKey:ID"`
	Name          string
	Cost          float32
	DeliverySpeed string
}
