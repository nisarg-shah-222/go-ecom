package models

import (
	"time"

	"gorm.io/datatypes"
)

type Product struct {
	Id          uint       `json:"id,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	ImageUrl    string     `json:"image_url"`
	Category    string     `json:"category"`
	Created     *time.Time `json:"-" gorm:"autoCreateTime"`
	Updated     *time.Time `json:"-" gorm:"autoUpdateTime"`
}

type ProductVersion struct {
	Id        uint           `json:"id,omitempty"`
	ProductId uint           `json:"product_id"`
	Details   datatypes.JSON `json:"details"`
	IsActive  bool           `json:"is_active"`
	Created   *time.Time     `json:"-" gorm:"autoCreateTime"`
	Updated   *time.Time     `json:"-" gorm:"autoUpdateTime"`
}
