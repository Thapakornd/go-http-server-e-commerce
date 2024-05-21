package models

import (
	"errors"
	"sync"
	"time"

	"gorm.io/gorm"
)

var mu sync.Mutex

type User struct {
	UserID      int64 `gorm:"primaryKey; autoIncremental:false"`
	FirstName   string
	LastName    string
	Email       string
	Username    string
	Password    string
	BirthOfDate time.Time
	Phone       string
	gorm.Model
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	mu.Lock()
	defer mu.Unlock()

	lastUser := User{}
	var lastIndex int

	err := db.Where("userID = ?", "1%").Last(&lastUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.UserID = int64(10000000) // 8 Digits include 1
		return nil
	}
	lastIndex = int(lastUser.UserID) + 1
	u.UserID = int64(lastIndex)
	return nil
}
