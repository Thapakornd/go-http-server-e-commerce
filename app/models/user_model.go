package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	IDS         int64 `gorm:"unique;autoIncremental:false"`
	FirstName   string
	LastName    string
	Email       string `gorm:"unique"`
	Username    string `gorm:"unique"`
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

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type GenerateToken struct {
	IDS      int64  `json:"ids" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type APIUser struct {
	IDS         uint64    `json:"ids"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	BirthOfDate time.Time `json:"birth_of_date"`
	Phone       string    `json:"phone"`
}

func (u *SignUpUser) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}
