package controller

import (
	"errors"
	"product/internal/constants"
	"product/internal/models"
	"product/internal/models/dtos"
)

func AddProduct(env *dtos.Env, addProductRequestDTOs dtos.AddProductRequestDTOs) error {
	filter := map[string]interface{}{
		"name": *addProductRequestDTOs.Name,
	}
	existingProductList := []models.Product{}
	env.MySQLConn.Table(constants.TableNameProduct).Where(filter).Find(&existingProductList)
	if len(existingProductList) > 0 {
		return errors.New("product already exists")
	}

	newProduct := models.Product{
		Name:        *addProductRequestDTOs.Name,
		Description: *addProductRequestDTOs.Description,
		Price:       *addProductRequestDTOs.Price,
		ImageUrl:    *addProductRequestDTOs.ImageUrl,
		Category:    *addProductRequestDTOs.Category,
	}

	newProductVersion := models.ProductVersion{
		Details:  *addProductRequestDTOs.Details,
		IsActive: true,
	}

	tx := env.MySQLConn.Begin()
	defer tx.Rollback()
	result := tx.Table(constants.TableNameProduct).Create(&newProduct)
	if result.Error != nil {
		return errors.Join(errors.New("unable to add product"), result.Error)
	}
	newProductVersion.ProductId = newProduct.Id
	result = tx.Table(constants.TableNameProductVersion).Create(&newProductVersion)
	if result.Error != nil {
		return errors.Join(errors.New("unable to add product version"), result.Error)
	}
	result = tx.Commit()
	if result.Error != nil {
		return errors.Join(errors.New("unable to commit to the database"), result.Error)
	}
	return nil
}
