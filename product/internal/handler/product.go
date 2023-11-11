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
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	var addProductRequestDTOs dtos.AddProductRequestDTOs
	if err := ctx.ShouldBindJSON(&addProductRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(addProductRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := controller.AddProduct(env, addProductRequestDTOs)
	if err != nil {
		env.Logger.Error(err.Error())
		if errors.Is(err, constants.ErrBadRequest) {
			response := constants.RESPONSE_BAD_REQUEST
			response["error"] = err.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, constants.RESPONSE_INTERNAL_SERVER_ERROR)
			return
		}
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
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	productIdParam := ctx.Param("id")
	productId, err := strconv.ParseInt(productIdParam, 10, 64)
	if err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = errors.New("path param not in the requested format").Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	getProductRequestDTOs := dtos.GetProductRequestDTOs{
		Id: &[]uint{uint(productId)}[0],
	}

	// Validate the request body
	if err := env.Validator.Struct(getProductRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	getProductResponseDTOs, err := controller.GetProduct(env, getProductRequestDTOs)
	if err != nil {
		if errors.Is(err, constants.ErrBadRequest) {
			env.Logger.Error(err.Error())
			response := constants.RESPONSE_BAD_REQUEST
			response["error"] = err.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		} else {
			env.Logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, constants.RESPONSE_INTERNAL_SERVER_ERROR)
			return
		}
	}

	// Add Product successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "success",
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
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	var updateProductRequestDTOs dtos.UpdateProductRequestDTOs
	if err := ctx.ShouldBindJSON(&updateProductRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productIdParam := ctx.Param("id")
	productId, err := strconv.ParseInt(productIdParam, 10, 64)
	if err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = errors.New("path param not in the requested format").Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	updateProductRequestDTOs.Id = &[]uint{uint(productId)}[0]

	// Validate the request body
	if err := env.Validator.Struct(updateProductRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if env.AuthDtos.Type != "X_API_KEY" && !controller.CheckUserPermission(env, *updateProductRequestDTOs.Id, "Update") {
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	err = controller.UpdateProduct(env, updateProductRequestDTOs)
	if err != nil {
		env.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, constants.RESPONSE_INTERNAL_SERVER_ERROR)
		return
	}

	// Add Product successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "successfully updated product",
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
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	var updateProductDetailsRequestDTOs dtos.UpdateProductDetailsRequestDTOs
	if err := ctx.ShouldBindJSON(&updateProductDetailsRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productIdParam := ctx.Param("id")
	productId, err := strconv.ParseInt(productIdParam, 10, 64)
	if err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = errors.New("path param not in the requested format").Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	updateProductDetailsRequestDTOs.Id = &[]uint{uint(productId)}[0]

	// Validate the request body
	if err := env.Validator.Struct(updateProductDetailsRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if env.AuthDtos.Type != "X_API_KEY" && !controller.CheckUserPermission(env, *updateProductDetailsRequestDTOs.Id, "Update") {
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	err = controller.UpdateProductDetails(env, updateProductDetailsRequestDTOs)
	if err != nil {
		env.Logger.Error(err.Error())
		if errors.Is(err, constants.ErrBadRequest) {
			response := constants.RESPONSE_BAD_REQUEST
			response["error"] = err.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, constants.RESPONSE_INTERNAL_SERVER_ERROR)
			return
		}
	}

	// Add Product successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "successfully updated product details",
		"data": map[string]interface{}{
			"success": true,
		},
	})
}

func DeleteProduct(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling delete product")

	if !utils.HasPermission(env, constants.DELETE_PRODUCT) {
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	productIdParam := ctx.Param("id")
	productId, err := strconv.ParseInt(productIdParam, 10, 64)
	if err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = errors.New("path param not in the requested format").Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	deleteProductRequestDTOs := dtos.DeleteProductRequestDTOs{
		Id: &[]uint{uint(productId)}[0],
	}

	if err := env.Validator.Struct(deleteProductRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if env.AuthDtos.Type != "X_API_KEY" && !controller.CheckUserPermission(env, *deleteProductRequestDTOs.Id, "Delete") {
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	err = controller.DeleteProduct(env, deleteProductRequestDTOs)
	if err != nil {
		env.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, constants.RESPONSE_INTERNAL_SERVER_ERROR)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "successfully deleted product",
	})
}
