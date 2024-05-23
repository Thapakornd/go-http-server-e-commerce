package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type OrderStore interface {
	GetAll() (*[]models.Order, error)
	GetByUserID(uint) (*models.Order, error)
	Create(*models.Order) error
	Update(*models.Order) error
	Delete(uint) error
}

type OrderStruct struct {
	db *gorm.DB
}

func NewOrderStore(db *gorm.DB) *OrderStruct {
	return &OrderStruct{
		db: db,
	}
}

func (as *OrderStruct) GetAll() (*[]models.Order, error) {
	return nil, nil
}

func (as *OrderStruct) GetByUserID(id uint) (*models.Order, error) {
	return nil, nil
}

func (as *OrderStruct) Create(o *models.Order) error {
	return nil
}

func (as *OrderStruct) Update(o *models.Order) error {
	return nil
}

func (as *OrderStruct) Delete(id uint) error {
	return nil
}
