package store

import (
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type UserStore interface {
	GetByID(uint) (*models.User, error)
	GetByEmail(string) (*models.User, error)
}

type UserStruct struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStruct {
	return &UserStruct{
		db: db,
	}
}

func (us *UserStruct) GetByID(id uint) (*models.User, error) {

	return nil, nil
}

func (us *UserStruct) GetByEmail(email string) (*models.User, error) {
	return nil, nil
}

func (us *UserStruct) GetByUsername(username string) (*models.User, error) {
	return nil, nil
}

func (us *UserStruct) Create(username string) error {
	return nil
}

func (us *UserStruct) Update(username string) (*models.User, error) {
	return nil, nil
}

func (us *UserStruct) Delete(username string) error {
	return nil
}
