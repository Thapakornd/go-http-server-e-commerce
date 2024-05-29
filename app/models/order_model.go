package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderItems       *[]OrderItem    `gorm:"foreignKey:OrderID"`
	ShippingMethod   *ShippingMethod `gorm:"foreignKey:ShippingMethodID"`
	PaymentMethod    *Payment        `gorm:"foreignKey:PaymentMethodID"`
	OrderStatus      OrderStatus     `gorm:"foreignKey:Status"`
	UserID           uint64
	BillingAddress   uint64
	PaymentMethodID  uint64
	ShippingMethodID uint64
	Status           uint64
}

type APIOrder struct {
	UserID           uint64
	BillingAddressID uint64
	PaymentID        uint64
	ShippingID       uint64
	Status           uint64
}
