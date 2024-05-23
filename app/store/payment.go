package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type PaymentStore interface {
	GetAll() (*[]models.Payment, error)
	GetByUserID(uint) (*[]models.Payment, error)
	Create(*models.Payment) error
	Update(*models.Payment) error
	Delete(uint) error
}

type PaymentStruct struct {
	db *gorm.DB
}

func NewPaymentStore(db *gorm.DB) *PaymentStruct {
	return &PaymentStruct{
		db: db,
	}
}

func (as *PaymentStruct) GetAll() (*[]models.Payment, error) {
	return nil, nil
}

func (as *PaymentStruct) GetByUserID(id uint) (*[]models.Payment, error) {
	return nil, nil
}

func (as *PaymentStruct) Create(p *models.Payment) error {
	return nil
}

func (as *PaymentStruct) Update(p *models.Payment) error {
	return nil
}

func (as *PaymentStruct) Delete(id uint) error {
	return nil
}
