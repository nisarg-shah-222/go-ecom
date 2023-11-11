package handler

import (
	"errors"
	"net/http"
	"order/internal/constants"
	"order/internal/controller"
	"order/internal/models/dtos"
	"order/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling create order")

	if !utils.HasPermission(env, constants.CREATE_ORDER) {
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	var createOrderRequestDTOs dtos.CreateOrderRequestDTOs
	if err := ctx.ShouldBindJSON(&createOrderRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(createOrderRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := controller.CreateOrder(env, createOrderRequestDTOs)
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

	// Create Order successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "order successful",
		"data": map[string]interface{}{
			"success": true,
		},
	})
}

func GetOrder(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling get order")

	if !utils.HasPermission(env, constants.GET_ORDER) {
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	orderIdParam := ctx.Param("id")
	orderId, err := strconv.ParseInt(orderIdParam, 10, 64)
	if err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = errors.New("path param not in the requested format").Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	getOrderRequestDTOs := dtos.GetOrderRequestDTOs{
		Id: &[]uint{uint(orderId)}[0],
	}

	// Validate the request body
	if err := env.Validator.Struct(getOrderRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	getOrderResponseDtos, err := controller.GetOrder(env, getOrderRequestDTOs)
	if err != nil {
		env.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, constants.RESPONSE_INTERNAL_SERVER_ERROR)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "success",
		"data":    getOrderResponseDtos,
	})
}
