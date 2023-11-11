package handler

import (
	"errors"
	"net/http"
	"strconv"
	"user/internal/constants"
	"user/internal/controller"
	"user/internal/models/dtos"
	"user/internal/utils"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// Retrieve the environment from the context
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	// Log the registration process
	env.Logger.Info("Handling registration")

	// Parse the request body into userRegisterRequestDTOs
	var userRegisterRequestDTOs dtos.UserRegisterRequestDTOs
	if err := ctx.ShouldBindJSON(&userRegisterRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(userRegisterRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Call the controller to register the user
	err := controller.Register(env, userRegisterRequestDTOs)
	if err != nil {
		env.Logger.Error(err.Error())
		if errors.Is(err, constants.ErrUserNameExists) {
			response := constants.RESPONSE_BAD_REQUEST
			response["error"] = err.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, constants.RESPONSE_INTERNAL_SERVER_ERROR)
			return
		}
	}

	// Registration successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "Registration successful",
	})
}

func Login(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling login")

	var userLoginRequestDTOs dtos.UserLoginRequestDTOs
	if err := ctx.ShouldBindJSON(&userLoginRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(userLoginRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Call the controller to register the user
	token, err := controller.Login(env, userLoginRequestDTOs)
	if err != nil {
		env.Logger.Error(err.Error())
		if errors.Is(err, constants.ErrUserNameNotExists) {
			response := constants.RESPONSE_BAD_REQUEST
			response["error"] = err.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, constants.RESPONSE_INTERNAL_SERVER_ERROR)
			return
		}
	}

	// Registration successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "login successful",
		"data": map[string]interface{}{
			"bearerToken": token,
		},
	})
}

func UpdateUser(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling update user")

	if !utils.HasPermission(env, constants.PATCH_USER) {
		ctx.JSON(http.StatusForbidden, constants.RESPONSE_FORBIDDEN)
		return
	}

	var updateUserRequestDTOs dtos.UpdateUserRequestDTOs
	if err := ctx.ShouldBindJSON(&updateUserRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userIdParam := ctx.Param("id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = errors.New("path param not in the requested format").Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	updateUserRequestDTOs.Id = &[]uint{uint(userId)}[0]

	// Validate the request body
	if err := env.Validator.Struct(updateUserRequestDTOs); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Call the controller to register the user
	err = controller.UpdateUser(env, updateUserRequestDTOs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, constants.RESPONSE_INTERNAL_SERVER_ERROR)
		return
	}

	// Registration successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "user updated successfully",
		"data": map[string]interface{}{
			"success": true,
		},
	})
}
