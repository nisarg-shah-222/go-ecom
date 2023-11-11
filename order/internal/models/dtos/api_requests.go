package dtos

import "gorm.io/datatypes"

type CreateOrderRequestDTOs struct {
	BillingAddress  *string                             `json:"billing_address" validate:"required"`
	ShippingAddress *string                             `json:"shipping_address" validate:"required"`
	TotalAmount     *float64                            `json:"total_amount" validate:"required"`
	ProductList     []CreateOrderProductListRequestDTOs `json:"product_list" validate:"required,min=1"`
}

type CreateOrderProductListRequestDTOs struct {
	Id       *uint    `json:"id" validate:"required"`
	Quantity *uint    `json:"quantity" validate:"required"`
	Price    *float64 `json:"price" validate:"required"`
}

type UserLoginRequestDTOs struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required,gt=7"`
}

type GetOrderRequestDTOs struct {
	Id *uint `json:"id" validate:"required"`
}

type GetOrderResponseDTOs struct {
	Id              uint                              `json:"id,omitempty"`
	UserId          uint                              `json:"order_id"`
	TotalAmount     float64                           `json:"total_amount"`
	BillingAddress  string                            `json:"billing_address"`
	ShippingAddress string                            `json:"shipping_address"`
	Status          string                            `json:"status"`
	ProductList     []GetOrderProductListResponseDTOs `json:"product_list"`
}

type GetOrderProductListResponseDTOs struct {
	Id          uint                            `json:"id"`
	ProductId   uint                            `json:"product_id"`
	Quantity    uint                            `json:"quantity"`
	Price       float64                         `json:"price"`
	ProductData GetOrderProductDataResponseDTOs `json:"product_data"`
}

type GetOrderProductDataResponseDTOs struct {
	Name        *string         `json:"name"`
	Description *string         `json:"description"`
	ImageUrl    *string         `json:"image_url"`
	Category    *string         `json:"category"`
	Details     *datatypes.JSON `json:"details"`
}
