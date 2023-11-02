package controller

import (
	"errors"
	"product/internal/constants"
	"product/internal/models"
	"product/internal/models/dtos"

	"gorm.io/gorm/clause"
)

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
		return errors.Join(errors.New("unable to update product"), result.Error)
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
	tx.Table(constants.TableNameProductVersion).Where(filter).Clauses(clause.Locking{Strength: "UPDATE"}).Find(&productVersionList)
	if len(productVersionList) == 0 {
		return errors.New("active product version not found")
	}
	filter = map[string]interface{}{
		"id": *updateProductDetailsRequestDTOs.VersionId,
	}
	updates := map[string]interface{}{
		"is_active": false,
	}
	result := tx.Table(constants.TableNameProductVersion).Where(filter).Updates(updates)
	if result.Error != nil {
		return errors.Join(errors.New("unable to update product version"), result.Error)
	}
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
