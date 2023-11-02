package dtos

import "gorm.io/datatypes"

type AddProductRequestDTOs struct {
	Name        *string         `json:"name" validate:"required"`
	Description *string         `json:"description" validate:"required"`
	Price       *float64        `json:"price" validate:"required"`
	ImageUrl    *string         `json:"image_url" validate:"required"`
	Category    *string         `json:"category" validate:"required"`
	Details     *datatypes.JSON `json:"details" validate:"required"`
}

type GetProductRequestDTOs struct {
	Id *uint `json:"id" validate:"required"`
}

type GetProductResponseDTOs struct {
	Id          uint           `json:"id,omitempty"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	ImageUrl    string         `json:"image_url"`
	Category    string         `json:"category"`
	VersionId   uint           `json:"version_id"`
	Details     datatypes.JSON `json:"details"`
}

type UpdateProductRequestDTOs struct {
	Id          *uint    `json:"id" validate:"required"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	ImageUrl    *string  `json:"image_url"`
	Category    *string  `json:"category"`
}

type UpdateProductDetailsRequestDTOs struct {
	Id        *uint           `json:"id" validate:"required"`
	VersionId *uint           `json:"version_id"`
	Details   *datatypes.JSON `json:"details"`
}
