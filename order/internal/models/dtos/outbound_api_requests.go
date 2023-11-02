package dtos

import "gorm.io/datatypes"

type GetProductDetailsResponseDTOs struct {
	Status  *uint           `json:"status"`
	Message *string         `json:"message"`
	Data    *ProductDetails `json:"data"`
}

type ProductDetails struct {
	Id          *uint           `json:"id" validate:"required"`
	Name        *string         `json:"name"`
	Description *string         `json:"description"`
	Price       *float64        `json:"price" validate:"required"`
	ImageUrl    *string         `json:"image_url"`
	Category    *string         `json:"category"`
	VersionId   *uint           `json:"version_id"`
	Details     *datatypes.JSON `json:"details"`
}
