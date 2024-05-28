package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type ProductStore interface {
	GetAll(int, int, *[]models.Product) (int64, error)
	GetByID(int) (*models.Product, error)
	Create(*models.APIProduct) error
	Update(*models.APIProduct, int) error
	Delete(int) error
}

type ProductStruct struct {
	db *gorm.DB
}

func NewProductStore(db *gorm.DB) *ProductStruct {
	return &ProductStruct{
		db: db,
	}
}

func (as *ProductStruct) GetAll(limit int, offset int, p *[]models.Product) (int64, error) {

	var total int64

	if err := as.db.Preload("Category").Limit(limit).Offset(offset).Order("id desc").Find(&p).Error; err != nil {
		return 0, err
	}

	if err := as.db.Table("categories").Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (as *ProductStruct) GetByID(id int) (*models.Product, error) {

	p := models.Product{}

	if err := as.db.Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

func (as *ProductStruct) Create(p *models.APIProduct) error {

	category_model := models.Category{}

	if err := as.db.Select("id").Where("name = ?", p.Category).First(&category_model).Error; err != nil {
		return err
	}

	new_product := models.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CategoryID:  uint64(category_model.ID),
	}

	if err := as.db.Create(&new_product).Error; err != nil {
		return err
	}

	return nil
}

func (as *ProductStruct) Update(p *models.APIProduct, id int) error {

	category_model := models.Category{}

	if err := as.db.Select("id").Where("name = ?", p.Category).First(&category_model).Error; err != nil {
		return err
	}

	update_product := models.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CategoryID:  uint64(category_model.ID),
	}

	if err := as.db.Model(&models.Product{}).Where("id = ?", id).Updates(&update_product).Error; err != nil {
		return err
	}

	return nil
}

func (as *ProductStruct) Delete(id int) error {

	if err := as.db.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return err
	}

	return nil
}
