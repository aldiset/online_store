package models

import "time"

type User struct {
	Id        uint      `gorm:"primary_key;" json:"id"`
	Fullname  string    `gorm:"size:255;not null;" json:"fullname"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null;" json:"password"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt  time.Time `gorm:"autoUpdateTime:false" json:"updated_at"`
}
