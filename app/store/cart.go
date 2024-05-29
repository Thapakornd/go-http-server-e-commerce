package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type CartStore interface {
	GetAll(int, int, *[]models.CartSession) (int64, error)
	GetByUserID(int, *models.CartSession) error
	GetByID(int, *models.CartSession) error
	Create(*models.CartSession) error
	Update(*models.CartSession, int) error
	Delete(int) error
}

type CartStruct struct {
	db *gorm.DB
}

func NewCartStore(db *gorm.DB) *CartStruct {
	return &CartStruct{
		db: db,
	}
}

func (as *CartStruct) GetAll(limit int, offset int, cs *[]models.CartSession) (int64, error) {

	var total int64

	if err := as.db.Preload("CartItem").Limit(limit).Offset(offset).Order("id desc").Find(&cs).Error; err != nil {
		return 0, nil
	}

	return total, nil
}

func (as *CartStruct) GetByUserID(id int, c *models.CartSession) error {

	if err := as.db.Where("created_by = ?", id).First(&c).Error; err != nil {
		return err
	}

	return nil
}

func (as *CartStruct) GetByID(id int, c *models.CartSession) error {

	if err := as.db.Where("id = ?", id).First(&c).Error; err != nil {
		return err
	}

	return nil
}

func (as *CartStruct) Create(cs *models.CartSession) error {

	if err := as.db.Create(&cs).Error; err != nil {
		return err
	}

	return nil
}

func (as *CartStruct) Update(cs *models.CartSession, id int) error {

	if err := as.db.Model(&models.CartSession{}).Where("id = ?", id).Updates(&cs).Error; err != nil {
		return err
	}

	return nil
}

func (as *CartStruct) Delete(id int) error {

	if err := as.db.Where("id = ?", id).Delete(&models.CartSession{}).Error; err != nil {
		return err
	}

	return nil
}
