package controller

import (
	"errors"
	"fmt"
	"order/internal/constants"
	"order/internal/models"
	"order/internal/models/dtos"
)

func CreateOrder(env *dtos.Env, createOrderRequestDTOs dtos.CreateOrderRequestDTOs) error {
	err := ValidateOrder(env, createOrderRequestDTOs)
	if err != nil {
		return errors.Join(constants.ErrValidationFailed, constants.ErrBadRequest, err)
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
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	for idx := range newOrderProductList {
		newOrderProductList[idx].OrderId = newOrder.Id
	}
	result = tx.Table(constants.TableNameOrderProduct).Create(newOrderProductList)
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	result = tx.Commit()
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	return nil
}

func ValidateOrder(env *dtos.Env, createOrderRequestDTOs dtos.CreateOrderRequestDTOs) error {
	if createOrderRequestDTOs.ProductList == nil || len(createOrderRequestDTOs.ProductList) == 0 {
		return errors.New("product list is empty")
	}
	totalAmount := float64(0)
	for _, product := range createOrderRequestDTOs.ProductList {
		productDetails, err := getProductDetails(env, *product.Id)
		if err != nil {
			return err
		}
		if *productDetails.Price != *product.Price {
			return errors.New(fmt.Sprintf("invalid product price %v", *product.Id))
		}
		totalAmount += *productDetails.Price * float64(*product.Quantity)
	}
	if totalAmount != *createOrderRequestDTOs.TotalAmount {
		return errors.New("invalid total price")
	}
	return nil
}

func GetOrder(env *dtos.Env, getOrderRequestDTOs dtos.GetOrderRequestDTOs) (*dtos.GetOrderResponseDTOs, error) {

	filter := map[string]interface{}{
		"id":      *getOrderRequestDTOs.Id,
		"user_id": *env.AuthDtos.Id,
	}
	existingOrderList := []models.Order{}
	env.MySQLConn.Table(constants.TableNameOrder).Where(filter).Find(&existingOrderList)
	if len(existingOrderList) == 0 {
		return nil, errors.Join(constants.ErrBadRequest, constants.ErrOrderNotExists)
	}
	existingOrder := existingOrderList[0]
	responseOrder := dtos.GetOrderResponseDTOs{
		Id:              existingOrder.Id,
		UserId:          existingOrder.UserId,
		TotalAmount:     existingOrder.TotalAmount,
		BillingAddress:  existingOrder.BillingAddress,
		ShippingAddress: existingOrder.ShippingAddress,
		Status:          existingOrder.Status,
	}

	filter = map[string]interface{}{
		"order_id": existingOrder.Id,
	}
	existingOrderProductList := []models.OrderProduct{}
	env.MySQLConn.Table(constants.TableNameOrderProduct).Where(filter).Find(&existingOrderProductList)
	responseProductList := []dtos.GetOrderProductListResponseDTOs{}
	for _, existingOrderProduct := range existingOrderProductList {
		responseProduct := dtos.GetOrderProductListResponseDTOs{
			Id:        existingOrderProduct.Id,
			ProductId: existingOrderProduct.ProductId,
			Quantity:  existingOrderProduct.Quantity,
			Price:     existingOrderProduct.Price,
		}
		productDetails, _ := getProductDetails(env, existingOrderProduct.ProductId)
		if productDetails != nil {
			responseProduct.ProductData = dtos.GetOrderProductDataResponseDTOs{
				Name:        productDetails.Name,
				Description: productDetails.Description,
				ImageUrl:    productDetails.ImageUrl,
				Category:    productDetails.Category,
				Details:     productDetails.Details,
			}
		}
		responseProductList = append(responseProductList, responseProduct)
	}
	responseOrder.ProductList = responseProductList
	return &responseOrder, nil
}
