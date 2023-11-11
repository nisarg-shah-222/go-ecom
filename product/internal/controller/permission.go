package controller

import (
	"errors"
	"product/internal/constants"
	"product/internal/models"
	"product/internal/models/dtos"

	"golang.org/x/exp/slices"
)

func AddProductUserPermission(env *dtos.Env, addProductUserPermissionRequestDTOs dtos.AddProductUserPermissionRequestDTOs) error {
	filter := map[string]interface{}{
		"product_id": *addProductUserPermissionRequestDTOs.ProductId,
		"user_id":    *addProductUserPermissionRequestDTOs.UserId,
	}
	existingProductUserPermissionList := []models.ProductUserPermission{}
	env.MySQLConn.Table(constants.TableNameProductUserPermission).Where(filter).Find(&existingProductUserPermissionList)
	if len(existingProductUserPermissionList) > 0 {
		existingProductUserPermission := existingProductUserPermissionList[0]
		if existingProductUserPermission.Permission == *addProductUserPermissionRequestDTOs.Permission {
			return errors.Join(constants.ErrBadRequest, constants.ErrPermissionExists)
		}
		updates := map[string]interface{}{
			"permission": *addProductUserPermissionRequestDTOs.Permission,
		}
		result := env.MySQLConn.Table(constants.TableNameProductUserPermission).Where(filter).Updates(updates)
		if result.Error != nil {
			return errors.Join(constants.ErrInternalServerError, result.Error)
		}
		return nil
	}
	newProductUserPermission := models.ProductUserPermission{
		ProductId:  *addProductUserPermissionRequestDTOs.ProductId,
		UserId:     *addProductUserPermissionRequestDTOs.UserId,
		Permission: *addProductUserPermissionRequestDTOs.Permission,
	}

	result := env.MySQLConn.Table(constants.TableNameProductUserPermission).Create(&newProductUserPermission)
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	return nil
}

func CheckUserPermission(env *dtos.Env, productId uint, requestedPermission string) bool {
	filter := map[string]interface{}{
		"product_id": productId,
		"user_id":    env.AuthDtos.Id,
	}
	existingProductUserPermissionList := []models.ProductUserPermission{}
	env.MySQLConn.Table(constants.TableNameProductUserPermission).Where(filter).Find(&existingProductUserPermissionList)
	if len(existingProductUserPermissionList) == 0 {
		return false
	}
	existingProductUserPermission := existingProductUserPermissionList[0]
	userPermissionList := constants.PermissionsHierarchyMap[existingProductUserPermission.Permission]
	return slices.Contains(userPermissionList, requestedPermission)
}
