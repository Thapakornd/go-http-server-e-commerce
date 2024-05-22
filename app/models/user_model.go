package models

import (
	"errors"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var mu sync.Mutex

type User struct {
	IDS         int64 `gorm:"autoIncremental:false"`
	FirstName   string
	LastName    string
	Email       string
	Username    string
	Password    string
	BirthOfDate time.Time
	Phone       string
	Addresses   *[]Address   `gorm:"foreignKey:UserID"`
	Payments    *[]Payment   `gorm:"foreignKey:UserID"`
	CartSession *CartSession `gorm:"foreignKey:CreatedBy"`
	Order       *Order       `gorm:"foreignKey:UserID"`
	gorm.Model
}

type SignInUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
}

type SignUpUser struct {
	FirstName   string    `json:"firstname" validate:"required"`
	LastName    string    `json:"lastname" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	BirthOfDate time.Time `json:"birth_of_date" validate:"required"`
	Phone       string    `json:"phone" validate:"required"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	mu.Lock()
	defer mu.Unlock()

	lastUser := User{}
	var lastIndex int

	err := db.Where("userID = ?", "1%").Last(&lastUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.IDS = int64(10000000) // 8 Digits include 1
		return nil
	}
	lastIndex = int(lastUser.IDS) + 1
	u.IDS = int64(lastIndex)
	return nil
}

func (u *SignUpUser) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}
