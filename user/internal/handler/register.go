package handler

import (
	"errors"
	"net/http"
	"user/internal/controller"
	"user/internal/models/dtos"

	"github.com/gin-gonic/gin"
)

// Register handles the registration of a user.
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("request not in the correct format"), err).Error(),
		})
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(userRegisterRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("validation error"), err).Error(),
		})
		return
	}

	// Call the controller to register the user
	err := controller.Register(env, userRegisterRequestDTOs)
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
		"message": "Registration successful",
	})
}
