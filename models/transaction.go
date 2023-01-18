package models

import "time"

type Transaction struct {
	Id                uint          `gorm:"primary_key;" json:"id"`
	CartID            int           `json:"cart_id"`
	PaymentMethodCode int           `json:"payment_method_code"`
	CreatedAt         time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt          time.Time     `gorm:"autoUpdateTime:true" json:"updated_at"`
	Cart              Cart          `gorm:"references:Id"`
	PaymentMethod     PaymentMethod `gorm:"references:Code"`
}
