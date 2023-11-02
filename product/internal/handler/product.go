package handler

import (
	"errors"
	"net/http"
	"product/internal/constants"
	"product/internal/controller"
	"product/internal/models/dtos"
	"product/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProduct(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling add product")

	if !utils.HasPermission(env, constants.ADD_PRODUCT) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  1,
			"message": "forbidden",
		})
		return
	}

	var addProductRequestDTOs dtos.AddProductRequestDTOs
	if err := ctx.ShouldBindJSON(&addProductRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("request not in the correct format"), err).Error(),
		})
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(addProductRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("validation error"), err).Error(),
		})
		return
	}

	err := controller.AddProduct(env, addProductRequestDTOs)
	if err != nil {
		if err.Error() == "product already exists" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  1,
				"message": "Bad request",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  1,
				"message": "Please try again after some time",
				"error":   err.Error(),
			})
		}
		return
	}

	// Add Product successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "successfully added product",
		"data": map[string]interface{}{
			"success": true,
		},
	})
}

func GetProduct(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling get product")

	if !utils.HasPermission(env, constants.GET_PRODUCT) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  1,
			"message": "forbidden",
		})
		return
	}

	productIdParam := ctx.Param("id")
	productId, err := strconv.ParseInt(productIdParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.New("path param not in the requested format"),
		})
		return
	}
	getProductRequestDTOs := dtos.GetProductRequestDTOs{
		Id: &[]uint{uint(productId)}[0],
	}

	// Validate the request body
	if err := env.Validator.Struct(getProductRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("validation error"), err).Error(),
		})
		return
	}

	getProductResponseDTOs, err := controller.GetProduct(env, getProductRequestDTOs)
	if err != nil {
		if err.Error() == "product does not exists" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  1,
				"message": "Bad request",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  1,
				"message": "Please try again after some time",
				"error":   err.Error(),
			})
		}
		return
	}

	// Add Product successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "successfully added product",
		"data":    getProductResponseDTOs,
	})
}

func UpdateProduct(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling update product")

	if !utils.HasPermission(env, constants.UPDATE_PRODUCT) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  1,
			"message": "forbidden",
		})
		return
	}

	var updateProductRequestDTOs dtos.UpdateProductRequestDTOs
	if err := ctx.ShouldBindJSON(&updateProductRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("request not in the correct format"), err).Error(),
		})
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(updateProductRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("validation error"), err).Error(),
		})
		return
	}

	err := controller.UpdateProduct(env, updateProductRequestDTOs)
	if err != nil {
		if err.Error() == "product already exists" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  1,
				"message": "Bad request",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  1,
				"message": "Please try again after some time",
				"error":   err.Error(),
			})
		}
		return
	}

	// Add Product successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "successfully added product",
		"data": map[string]interface{}{
			"success": true,
		},
	})
}

func UpdateProductDetails(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling update product details")

	if !utils.HasPermission(env, constants.UPDATE_PRODUCT) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  1,
			"message": "forbidden",
		})
		return
	}

	var updateProductDetailsRequestDTOs dtos.UpdateProductDetailsRequestDTOs
	if err := ctx.ShouldBindJSON(&updateProductDetailsRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("request not in the correct format"), err).Error(),
		})
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(updateProductDetailsRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("validation error"), err).Error(),
		})
		return
	}

	err := controller.UpdateProductDetails(env, updateProductDetailsRequestDTOs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  1,
			"message": "Please try again after some time",
			"error":   err.Error(),
		})
		return
	}

	// Add Product successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "successfully added product",
		"data": map[string]interface{}{
			"success": true,
		},
	})
}
