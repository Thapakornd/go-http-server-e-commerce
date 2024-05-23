package store

import (
	"errors"
	"fmt"
	"sync"

	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

var mu sync.Mutex

type UserStore interface {
	GetByID(uint) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	GetByUsername(string) (*models.User, error)
	Create(*models.User) error
	Update(*models.User) error
	Delete(uint) error
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

func (us *UserStruct) Create(u *models.User) error {
	mu.Lock()
	defer mu.Unlock()

	var lastIndex models.User

	err := us.db.Last(&lastIndex).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.IDS = 10000000
		} else {
			return err
		}
	} else {
		u.IDS = lastIndex.IDS + 1
	}

	if err := us.db.Create(&u); err.Error != nil {
		fmt.Println(err.Error.Error())
		return err.Error
	}
	return nil
}

func (us *UserStruct) Update(u *models.User) error {
	return nil
}

func (us *UserStruct) Delete(id uint) error {
	return nil
}
