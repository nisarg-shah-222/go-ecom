package controller

import (
	"errors"
	"fmt"
	"order/internal/models/dtos"
)

func ValidateOrder(env *dtos.Env, createOrderRequestDTOs dtos.CreateOrderRequestDTOs) error {
	if createOrderRequestDTOs.ProductList == nil || len(createOrderRequestDTOs.ProductList) == 0 {
		return errors.New("product list is empty")
	}
	totalAmount := float64(0)
	for _, product := range createOrderRequestDTOs.ProductList {
		productPrice, err := getProductPrice(env, *product.Id)
		if err != nil {
			return err
		}
		if productPrice != *product.Price {
			return errors.New(fmt.Sprintf("invalid product price %v", product.Id))
		}
		totalAmount += productPrice * float64(*product.Quantity)
	}
	if totalAmount != *createOrderRequestDTOs.TotalAmount {
		return errors.New("invalid total price")
	}
	return nil
}
