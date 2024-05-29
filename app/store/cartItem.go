package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type CartItemStore interface {
	GetAll(int, int, *[]models.CartItem) (int64, error)
	GetByUserID(int) (*models.CartItem, error)
	GetByProductID(int, int) (*models.CartItem, error)
	Create(*models.CartItem) error
	Update(*models.CartItem, int) error
	Delete(int) error
}

type CartItemStruct struct {
	db *gorm.DB
}

func NewCartItemStore(db *gorm.DB) *CartItemStruct {
	return &CartItemStruct{
		db: db,
	}
}

func (ci *CartItemStruct) GetAll(limit int, offset int, c *[]models.CartItem) (int64, error) {
	return 0, nil
}

func (ci *CartItemStruct) GetByProductID(id int, session_id int) (*models.CartItem, error) {

	c := models.CartItem{}

	if err := ci.db.Where("product_id = ? AND cart_session_id = ?", id, session_id).Find(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func (ci *CartItemStruct) GetByUserID(id int) (*models.CartItem, error) {
	return nil, nil
}

func (ci *CartItemStruct) Create(c *models.CartItem) error {

	if err := ci.db.Create(&c).Error; err != nil {
		return err
	}

	return nil
}

func (ci *CartItemStruct) Update(c *models.CartItem, id int) error {

	if err := ci.db.Model(&models.CartItem{}).Where("id = ?", id).Updates(c).Error; err != nil {
		return err
	}

	return nil
}

func (ci *CartItemStruct) Delete(id int) error {
	return nil
}
