package models

import "time"

type OrderProduct struct {
	Id        uint       `json:"id,omitempty"`
	OrderId   uint       `json:"order_id"`
	ProductId uint       `json:"product_id"`
	Quantity  uint       `json:"quantity"`
	Price     float64    `json:"price"`
	Created   *time.Time `json:"-" gorm:"autoCreateTime"`
	Updated   *time.Time `json:"-" gorm:"autoUpdateTime"`
}
