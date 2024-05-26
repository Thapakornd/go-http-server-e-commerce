package store

import (
	"fmt"

	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type ProductStore interface {
	GetAll(int, int, *[]models.APIProduct) error
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

func (as *ProductStruct) GetAll(limit int, offset int, p *[]models.APIProduct) error {

	// var total int64
	p_model := &[]models.Product{}

	if err := as.db.Preload("Category").Find(&p_model).Error; err != nil {
		return err
	}

	fmt.Println(p_model)

	return nil
}

func (as *ProductStruct) GetByID(id uint) (*models.Product, error) {
	return nil, nil
}

func (as *ProductStruct) Create(p *models.Product) error {

	if err := as.db.Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func (as *ProductStruct) Update(p *models.Product) error {
	return nil
}

func (as *ProductStruct) Delete(id uint) error {
	return nil
}
