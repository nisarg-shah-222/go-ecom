package server

import (
	"fmt"
	"order/internal/constants"
	"order/internal/handler"
	"order/internal/server/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Initialize() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.Use(gin.CustomRecovery(middleware.ErrorHandler))
	route(&router.RouterGroup)
	router.Run(fmt.Sprintf(":%s", viper.GetString(constants.PORT)))
}

func route(router *gin.RouterGroup) {
	router.Use(middleware.EnvMiddleware)
	public := router.Group("/")
	router.Use(middleware.AuthMiddleware)
	router.Use(middleware.RateLimiterMiddleware)
	private := router.Group("/")
	privateV1Apis := private.Group("/api/v1")
	{
		privateV1Apis.POST("/order", handler.CreateOrder)
		privateV1Apis.GET("/order/:id", handler.GetOrder)
	}
	public.GET("/health-check", handler.Healthcheck)
}
