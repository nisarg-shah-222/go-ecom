package controller

import (
	"errors"
	"order/internal/constants"
	"order/internal/models"
	"order/internal/models/dtos"
)

func CreateOrder(env *dtos.Env, createOrderRequestDTOs dtos.CreateOrderRequestDTOs) error {
	err := ValidateOrder(env, createOrderRequestDTOs)
	if err != nil {
		return errors.Join(errors.New("order validation failed"), err)
	}

	newOrder := models.Order{
		UserId:          *env.AuthDtos.Id,
		TotalAmount:     *createOrderRequestDTOs.TotalAmount,
		BillingAddress:  *createOrderRequestDTOs.BillingAddress,
		ShippingAddress: *createOrderRequestDTOs.ShippingAddress,
		Status:          "CREATED",
	}

	newOrderProductList := []models.OrderProduct{}
	for _, product := range createOrderRequestDTOs.ProductList {
		newOrderProduct := models.OrderProduct{
			ProductId: *product.Id,
			Quantity:  *product.Quantity,
			Price:     *product.Price,
		}
		newOrderProductList = append(newOrderProductList, newOrderProduct)
	}

	tx := env.MySQLConn.Begin()
	defer tx.Rollback()
	result := tx.Table(constants.TableNameOrder).Create(&newOrder)
	if result.Error != nil {
		return errors.Join(errors.New("unable to create order"), result.Error)
	}
	for idx := range newOrderProductList {
		newOrderProductList[idx].OrderId = newOrder.Id
	}
	result = tx.Table(constants.TableNameOrderProduct).Create(newOrderProductList)
	if result.Error != nil {
		return errors.Join(errors.New("unable to add order products"), result.Error)
	}
	result = tx.Commit()
	if result.Error != nil {
		return errors.Join(errors.New("unable to commit to the database"), result.Error)
	}
	return nil
}
