package handler

import (
	"errors"
	"net/http"
	"order/internal/constants"
	"order/internal/controller"
	"order/internal/models/dtos"
	"order/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling login")

	if !utils.HasPermission(env, constants.CREATE_ORDER) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  1,
			"message": "Forbidden",
		})
		return
	}

	var createOrderRequestDTOs dtos.CreateOrderRequestDTOs
	if err := ctx.ShouldBindJSON(&createOrderRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("request not in the correct format"), err).Error(),
		})
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(createOrderRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("validation error"), err).Error(),
		})
		return
	}

	err := controller.CreateOrder(env, createOrderRequestDTOs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  1,
			"message": "Please try again after some time",
			"error":   err.Error(),
		})
		return
	}

	// Create Order successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "order successful",
		"data": map[string]interface{}{
			"success": true,
		},
	})
}
