package handler

import (
	"errors"
	"net/http"
	"user/internal/controller"
	"user/internal/models/dtos"

	"github.com/gin-gonic/gin"
)

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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("request not in the correct format"), err).Error(),
		})
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(userLoginRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("validation error"), err).Error(),
		})
		return
	}

	// Call the controller to register the user
	token, err := controller.Login(env, userLoginRequestDTOs)
	if err != nil {
		if err.Error() == "username already exists" {
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

	// Registration successful
	ctx.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "login successful",
		"data": map[string]interface{}{
			"bearerToken": token,
		},
	})
}
