package handler

import (
	"errors"
	"net/http"
	"product/internal/constants"
	"product/internal/controller"
	"product/internal/models/dtos"
	"product/internal/utils"

	"github.com/gin-gonic/gin"
)

func AddProductUserPermission(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling add product user permission")

	if !utils.HasPermission(env, constants.ADD_PRODUCT_USER_PERMISSION) {
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	var addProductUserPermissionRequestDTOs dtos.AddProductUserPermissionRequestDTOs
	if err := ctx.ShouldBindJSON(&addProductUserPermissionRequestDTOs); err != nil {
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(addProductUserPermissionRequestDTOs); err != nil {
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := controller.AddProductUserPermission(env, addProductUserPermissionRequestDTOs)
	if err != nil {
		if errors.Is(err, constants.ErrPermissionExists) {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  0,
				"message": "permission already exists",
			})
			return
		}

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
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  0,
		"message": "successfully added permission",
		"data": map[string]interface{}{
			"success": true,
		},
	})
}
