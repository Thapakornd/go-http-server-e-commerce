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
	GetAll(int, int, *[]models.APIUser) (int64, error)
	GetByID(uint) (*models.APIUser, error)
	GetByEmail(string) (*models.User, error)
	GetByUsername(string) (*models.User, error)
	Create(*models.User) error
	Update(*models.APIUser, int) error
	Delete(int) error
}

type UserStruct struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStruct {
	return &UserStruct{
		db: db,
	}
}

func (us *UserStruct) GetAll(limit int, offset int, u *[]models.APIUser) (int64, error) {

	var total int64

	if err := us.db.Model(&models.User{}).Limit(limit).Offset(offset).Order("id_s desc").Find(&u).Error; err != nil {
		return 0, err
	}

	if err := us.db.Table("users").Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (us *UserStruct) GetByID(id uint) (*models.APIUser, error) {

	u := &models.APIUser{}

	if err := us.db.Model(&models.User{}).Where("id_s = ?", id).First(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (us *UserStruct) GetByEmail(email string) (*models.User, error) {

	u := &models.User{}

	if err := us.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (us *UserStruct) GetByUsername(username string) (*models.User, error) {

	u := &models.User{}

	if err := us.db.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
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

func (us *UserStruct) Update(u *models.APIUser, id int) error {

	if err := us.db.Model(&models.User{}).Where("id_s = ?", id).Updates(&u).Error; err != nil {
		return err
	}

	return nil
}

func (us *UserStruct) Delete(id int) error {

	if err := us.db.Where("id_s = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}
