package models

import "time"

type Transaction struct {
	Id                uint          `gorm:"primary_key;" json:"id"`
	CartID            uint          `json:"cart_id"`
	PaymentMethodCode string        `json:"payment_method_code"`
	CreatedAt         time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt          time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Cart              Cart          `gorm:"references:Id"`
	PaymentMethod     PaymentMethod `gorm:"references:Code"`
}
