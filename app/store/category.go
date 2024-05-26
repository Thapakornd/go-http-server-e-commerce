package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type CategoryStore interface {
	GetAllCategories(*[]models.Category) error
	GetCategoryByID(uint) error
	Create(*models.Category) error
	Update(*models.Category, uint) error
	Delete(uint) error
}

type CategoryStruct struct {
	db *gorm.DB
}

func NewCategoryStore(db *gorm.DB) *CategoryStruct {
	return &CategoryStruct{
		db: db,
	}
}

func (css *CategoryStruct) GetAllCategories(cs *[]models.Category) error {

	return nil
}

func (css *CategoryStruct) GetCategoryByID(id uint) error {
	return nil
}

func (css *CategoryStruct) Create(cs *models.Category) error {
	return nil
}

func (css *CategoryStruct) Update(cs *models.Category, id uint) error {
	return nil
}

func (css *CategoryStruct) Delete(id uint) error {
	return nil
}
