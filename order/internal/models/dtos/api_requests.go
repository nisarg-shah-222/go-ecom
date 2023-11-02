package dtos

type CreateOrderRequestDTOs struct {
	BillingAddress  *string                             `json:"billing_address" validate:"required"`
	ShippingAddress *string                             `json:"shipping_address" validate:"required"`
	TotalAmount     *float64                            `json:"total_amount" validate:"required"`
	ProductList     []CreateOrderProductListRequestDTOs `json:"product_list" validate:"required,len=1"`
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
