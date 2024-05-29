package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	User        User `gorm:"foreignKey:UserID"`
	UserID      int64
	Village     string
	Road        string
	SubDistrict string
	District    string
	Province    string
	PostalCode  string `gorm:"type:char(5)"`
}

type APIAddress struct {
	IDS         int64  `json:"id_user" validate:"required"`
	Village     string `json:"village" validate:"required"`
	Road        string `json:"road" validate:"required"`
	SubDistrict string `json:"sub_district" validate:"required"`
	District    string `json:"district" validate:"required"`
	Province    string `json:"province" validate:"required"`
	PostalCode  string `json:"postal_code" validate:"required"`
}
