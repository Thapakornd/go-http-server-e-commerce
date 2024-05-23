package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type CartStore interface {
	GetAll() (*[]models.CartSession, error)
	GetByID(uint) (*models.CartSession, error)
	Create(*models.CartSession) error
	Update(*models.CartSession) error
	Delete(uint) error
}

type CartStruct struct {
	db *gorm.DB
}

func NewCartStore(db *gorm.DB) *CartStruct {
	return &CartStruct{
		db: db,
	}
}

func (as *CartStruct) GetAll() (*[]models.CartSession, error) {
	return nil, nil
}

func (as *CartStruct) GetByID(id uint) (*models.CartSession, error) {
	return nil, nil
}

func (as *CartStruct) Create(*models.CartSession) error {
	return nil
}

func (as *CartStruct) Update(*models.CartSession) error {
	return nil
}

func (as *CartStruct) Delete(id uint) error {
	return nil
}
