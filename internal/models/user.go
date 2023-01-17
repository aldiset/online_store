package models

import (
	"time"
)

type User struct {
	Id          uint      `gorm:"primary_key;" json:"id"`
	Fullname    string    `gorm:"size:255;not null;" json:"fullname"`
	Username    string    `gorm:"size:255;not null;unique" json:"username"`
	Password    string    `gorm:"size:255;not null;" json:"password"`
	Email       string    `gorm:"size:255;not null;unique" json:"email"`
	CreatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_date"`
}

func (u *User) Create() (*User, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
