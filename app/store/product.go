package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type ProductStore interface {
	GetAll() (*[]models.Product, error)
	GetByID(uint) (*models.Product, error)
	Create(*models.Product) error
	Update(*models.Product) error
	Delete(uint) error
}

type ProductStruct struct {
	db *gorm.DB
}

func NewProductStore(db *gorm.DB) *ProductStruct {
	return &ProductStruct{
		db: db,
	}
}

func (as *ProductStruct) GetAll() (*[]models.Product, error) {
	return nil, nil
}

func (as *ProductStruct) GetByID(id uint) (*models.Product, error) {
	return nil, nil
}

func (as *ProductStruct) Create(p *models.Product) error {
	return nil
}

func (as *ProductStruct) Update(p *models.Product) error {
	return nil
}

func (as *ProductStruct) Delete(id uint) error {
	return nil
}
