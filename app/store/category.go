package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type CategoryStore interface {
	GetAllCategories(*[]models.Category) error
	GetCategoryByID(uint) (*models.Category, error)
	Create(*models.Category) error
	Update(*models.Category, int) error
	Delete(int) error
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

	if err := css.db.Find(&cs).Error; err != nil {
		return err
	}

	return nil
}

func (css *CategoryStruct) GetCategoryByID(id uint) (*models.Category, error) {

	category := &models.Category{}

	if err := css.db.Where("id = ?", id).Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (css *CategoryStruct) Create(cs *models.Category) error {

	if err := css.db.Create(&cs).Error; err != nil {
		return err
	}

	return nil
}

func (css *CategoryStruct) Update(cs *models.Category, id int) error {

	if err := css.db.Model(&models.Category{}).Where("id = ?", id).Updates(&cs).Error; err != nil {
		return err
	}

	return nil
}

func (css *CategoryStruct) Delete(id int) error {

	if err := css.db.Where("id = ?", id).Delete(&models.Category{}).Error; err != nil {
		return err
	}

	return nil
}
