package models

import (
	"time"
)

type Category struct {
	Id        uint      `gorm:"primary_key;" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Code      string    `gorm:"size:255;not null;unique" json:"code"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
