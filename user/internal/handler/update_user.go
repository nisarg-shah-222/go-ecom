package handler

import (
	"errors"
	"net/http"
	"user/internal/constants"
	"user/internal/controller"
	"user/internal/models/dtos"
	"user/internal/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUser(ctx *gin.Context) {
	val, ok := ctx.Get("env")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	env := val.(*dtos.Env)
	env.Logger.Info("Handling login")

	if !utils.HasPermission(env, constants.PATCH_USER) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  1,
			"message": "forbidden",
		})
		return
	}

	var updateUserRequestDTOs dtos.UpdateUserRequestDTOs
	if err := ctx.ShouldBindJSON(&updateUserRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("request not in the correct format"), err).Error(),
		})
		return
	}

	// Validate the request body
	if err := env.Validator.Struct(updateUserRequestDTOs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Bad request",
			"error":   errors.Join(errors.New("validation error"), err).Error(),
		})
		return
	}

	// Call the controller to register the user
	err := controller.UpdateUser(env, updateUserRequestDTOs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  1,
			"message": "Please try again after some time",
			"error":   err.Error(),
		})
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
