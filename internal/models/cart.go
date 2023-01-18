package models

import "time"

type Cart struct {
	Id        uint      `gorm:"primary_key;" json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt  time.Time `gorm:"autoUpdateTime:true" json:"updated_at"`
	User      User      `gorm:"references:Id"` //use User.Id
	Product   Product   `gorm:"references:Id"` //use Product.Id
}
