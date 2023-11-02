package handler

import (
	"net/http"
	"order/internal/models/dtos"

	"github.com/gin-gonic/gin"
)

func Healthcheck(ctx *gin.Context) {
	var env *dtos.Env
	if lo, ok := ctx.Get("env"); ok {
		if val, ok := lo.(*dtos.Env); ok {
			env = val
		} else {
			ctx.AbortWithStatus(500)
			return
		}
	} else {
		ctx.AbortWithStatus(500)
		return
	}
	env.Logger.Info("handling health-check")
	ctx.JSON(http.StatusOK, gin.H{"status": 0, "message": "service is up"})
}
