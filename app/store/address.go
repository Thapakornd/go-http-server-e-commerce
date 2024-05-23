package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type AddressStore interface {
	GetAll() (*[]models.Address, error)
	GetByUserID(uint) (*models.Address, error)
	Create(*models.Address) error
	Update(*models.Address) error
	Delete(uint) error
}

type AddressStruct struct {
	db *gorm.DB
}

func NewAddressStore(db *gorm.DB) *AddressStruct {
	return &AddressStruct{
		db: db,
	}
}

func (as *AddressStruct) GetAll() (*[]models.Address, error) {
	return nil, nil
}

func (as *AddressStruct) GetByUserID(id uint) (*models.Address, error) {
	return nil, nil
}

func (as *AddressStruct) Create(a *models.Address) error {
	return nil
}

func (as *AddressStruct) Update(a *models.Address) error {
	return nil
}

func (as *AddressStruct) Delete(id uint) error {
	return nil
}
