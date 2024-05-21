package models

import (
	"errors"
	"sync"
	"time"

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
