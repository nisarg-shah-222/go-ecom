package controller

import (
	"errors"
	"product/internal/constants"
	"product/internal/models"
	"product/internal/models/dtos"
)

func GetProduct(env *dtos.Env, getProductRequestDTOs dtos.GetProductRequestDTOs) (*dtos.GetProductResponseDTOs, error) {
	filter := map[string]interface{}{
		"id": *getProductRequestDTOs.Id,
	}
	existingProductList := []models.Product{}
	env.MySQLConn.Table(constants.TableNameProduct).Where(filter).Find(&existingProductList)
	if len(existingProductList) == 0 {
		return nil, errors.New("product does not exists")
	}
	existingProduct := existingProductList[0]

	filter = map[string]interface{}{
		"product_id": *getProductRequestDTOs.Id,
		"is_active":  true,
	}
	existingProductVersionList := []models.ProductVersion{}
	env.MySQLConn.Table(constants.TableNameProductVersion).Where(filter).Find(&existingProductVersionList)
	if len(existingProductVersionList) == 0 {
		return nil, errors.New("product details does not exist")
	}
	existingProductVersion := existingProductVersionList[0]
	getProductResponseDTOs := dtos.GetProductResponseDTOs{
		Id:          existingProduct.Id,
		Name:        existingProduct.Name,
		Description: existingProduct.Description,
		Price:       existingProduct.Price,
		ImageUrl:    existingProduct.ImageUrl,
		Category:    existingProduct.Category,
		VersionId:   existingProductVersion.Id,
		Details:     existingProductVersion.Details,
	}
	return &getProductResponseDTOs, nil
}
