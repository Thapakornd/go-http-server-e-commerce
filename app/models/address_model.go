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
