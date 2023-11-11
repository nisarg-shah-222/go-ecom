package controller

import (
	"errors"
	"product/internal/constants"
	"product/internal/models"
	"product/internal/models/dtos"

	"gorm.io/gorm/clause"
)

func AddProduct(env *dtos.Env, addProductRequestDTOs dtos.AddProductRequestDTOs) error {
	filter := map[string]interface{}{
		"name": *addProductRequestDTOs.Name,
	}
	existingProductList := []models.Product{}
	env.MySQLConn.Table(constants.TableNameProduct).Where(filter).Find(&existingProductList)
	if len(existingProductList) > 0 {
		return errors.Join(constants.ErrBadRequest, constants.ErrProductExists)
	}

	newProduct := models.Product{
		Name:        *addProductRequestDTOs.Name,
		Description: *addProductRequestDTOs.Description,
		Price:       *addProductRequestDTOs.Price,
		ImageUrl:    *addProductRequestDTOs.ImageUrl,
		Category:    *addProductRequestDTOs.Category,
		IsActive:    true,
	}

	newProductVersion := models.ProductVersion{
		Details:  *addProductRequestDTOs.Details,
		IsActive: true,
	}

	newProductUserPermission := models.ProductUserPermission{
		UserId:     *env.AuthDtos.Id,
		Permission: "Delete",
	}

	tx := env.MySQLConn.Begin()
	defer tx.Rollback()
	result := tx.Table(constants.TableNameProduct).Create(&newProduct)
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	newProductVersion.ProductId = newProduct.Id
	result = tx.Table(constants.TableNameProductVersion).Create(&newProductVersion)
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	newProductUserPermission.ProductId = newProduct.Id
	result = tx.Table(constants.TableNameProductUserPermission).Create(&newProductUserPermission)
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	result = tx.Commit()
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	return nil
}

func GetProduct(env *dtos.Env, getProductRequestDTOs dtos.GetProductRequestDTOs) (*dtos.GetProductResponseDTOs, error) {
	filter := map[string]interface{}{
		"id": *getProductRequestDTOs.Id,
	}
	existingProductList := []models.Product{}
	env.MySQLConn.Table(constants.TableNameProduct).Where(filter).Find(&existingProductList)
	if len(existingProductList) == 0 {
		return nil, errors.Join(constants.ErrBadRequest, constants.ErrProductNotExists)
	}
	existingProduct := existingProductList[0]

	filter = map[string]interface{}{
		"product_id": *getProductRequestDTOs.Id,
		"is_active":  true,
	}
	existingProductVersionList := []models.ProductVersion{}
	env.MySQLConn.Table(constants.TableNameProductVersion).Where(filter).Find(&existingProductVersionList)
	if len(existingProductVersionList) == 0 {
		return nil, errors.Join(constants.ErrBadRequest, constants.ErrProductDetailNotExists)
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

func UpdateProduct(env *dtos.Env, updateProductRequestDTOs dtos.UpdateProductRequestDTOs) error {
	filter := map[string]interface{}{
		"id": *updateProductRequestDTOs.Id,
	}
	updates := map[string]interface{}{}
	if updateProductRequestDTOs.Category != nil {
		updates["category"] = updateProductRequestDTOs.Category
	}
	if updateProductRequestDTOs.Description != nil {
		updates["description"] = updateProductRequestDTOs.Description
	}
	if updateProductRequestDTOs.ImageUrl != nil {
		updates["image_url"] = updateProductRequestDTOs.ImageUrl
	}
	if updateProductRequestDTOs.Price != nil {
		updates["price"] = updateProductRequestDTOs.Price
	}
	if updateProductRequestDTOs.Name != nil {
		updates["name"] = updateProductRequestDTOs.Name
	}
	result := env.MySQLConn.Table(constants.TableNameProduct).Where(filter).Updates(updates)
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	return nil
}

func UpdateProductDetails(env *dtos.Env, updateProductDetailsRequestDTOs dtos.UpdateProductDetailsRequestDTOs) error {
	filter := map[string]interface{}{
		"id":        *updateProductDetailsRequestDTOs.VersionId,
		"is_active": true,
	}

	newProductVersion := models.ProductVersion{
		ProductId: *updateProductDetailsRequestDTOs.Id,
		Details:   *updateProductDetailsRequestDTOs.Details,
		IsActive:  true,
	}

	tx := env.MySQLConn.Begin()
	defer tx.Rollback()
	productVersionList := []models.ProductVersion{}
	result := tx.Debug().Table(constants.TableNameProductVersion).Where(filter).Clauses(clause.Locking{Strength: "UPDATE"}).Find(&productVersionList)
	if result.Error != nil {
		return result.Error
	}
	if len(productVersionList) == 0 {
		return errors.Join(constants.ErrBadRequest, constants.ErrProductDetailNotExists)
	}
	filter = map[string]interface{}{
		"id": *updateProductDetailsRequestDTOs.VersionId,
	}
	updates := map[string]interface{}{
		"is_active": false,
	}
	result = tx.Table(constants.TableNameProductVersion).Where(filter).Updates(updates)
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	result = tx.Table(constants.TableNameProductVersion).Create(&newProductVersion)
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	result = tx.Commit()
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	return nil
}

func DeleteProduct(env *dtos.Env, deleteProductRequestDTOs dtos.DeleteProductRequestDTOs) error {
	filter := map[string]interface{}{
		"id": *deleteProductRequestDTOs.Id,
	}
	existingProductList := []models.Product{}
	env.MySQLConn.Table(constants.TableNameProduct).Where(filter).Find(&existingProductList)
	if len(existingProductList) == 0 {
		return errors.Join(constants.ErrBadRequest, constants.ErrProductNotExists)
	}
	existingProduct := existingProductList[0]
	if !existingProduct.IsActive {
		return errors.Join(constants.ErrBadRequest, constants.ErrProductNotExists)
	}

	tx := env.MySQLConn.Begin()
	defer tx.Rollback()

	result := tx.Table(constants.TableNameProduct).Where(filter).Update("is_active", false)
	if result.Error != nil {
		errors.Join(constants.ErrInternalServerError, result.Error)
	}
	filter = map[string]interface{}{
		"product_id": *deleteProductRequestDTOs.Id,
	}
	result = tx.Table(constants.TableNameProductVersion).Where(filter).Update("is_active", false)
	if result.Error != nil {
		errors.Join(constants.ErrInternalServerError, result.Error)
	}
	result = tx.Commit()
	if result.Error != nil {
		errors.Join(constants.ErrInternalServerError, result.Error)
	}
	return nil
}
