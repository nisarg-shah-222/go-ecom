package middleware

import (
	"fmt"
	"net/http"
	"order/internal/models/dtos"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

func RateLimiterMiddleware(ctx *gin.Context) {
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
	if env.AuthDtos.Type == "AUTH_TOKEN" {
		res, err := env.RateLimiter.Allow(ctx, fmt.Sprintf("%v:%v", *env.AuthDtos.Id, ctx.Request.URL.Path), redis_rate.PerMinute(5))
		if err != nil {
			panic(err)
		}
		if res.Allowed == 0 {
			AbortWith429(ctx)
		}
	}
	ctx.Next()
}

func AbortWith429(ctx *gin.Context) {
	ctx.JSON(http.StatusTooManyRequests, gin.H{"status": 1, "message": "too many requests, try again after some time"})
	ctx.AbortWithStatus(http.StatusTooManyRequests)
}
