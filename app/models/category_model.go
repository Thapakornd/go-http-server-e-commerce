package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `json:"category_name" validate:"required" gorm:"unique;"`
	Products *[]Product
}
