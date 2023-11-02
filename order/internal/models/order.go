package models

import "time"

type Order struct {
	Id              uint       `json:"id,omitempty"`
	UserId          uint       `json:"order_id"`
	TotalAmount     float64    `json:"total_amount"`
	BillingAddress  string     `json:"billing_address"`
	ShippingAddress string     `json:"shipping_address"`
	Status          string     `json:"status"`
	Created         *time.Time `json:"-" gorm:"autoCreateTime"`
	Updated         *time.Time `json:"-" gorm:"autoUpdateTime"`
}
