package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Product   Product `gorm:"foreignKey:ProductID"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	ProductID uint64
	OrderID   uint64
	Quantity  int
}
