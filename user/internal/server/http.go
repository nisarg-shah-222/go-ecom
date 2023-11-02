package server

import (
	"fmt"
	"user/internal/constants"
	"user/internal/handler"
	"user/internal/server/middleware"

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
	private := router.Group("/")
	privateV1Apis := private.Group("/api/v1")
	{
		privateV1Apis.PATCH("/user", handler.UpdateUser)
	}
	publicV1Apis := public.Group("/api/v1")
	{
		publicV1Apis.POST("/login", handler.Login)
		publicV1Apis.POST("/register", handler.Register)
	}
	public.GET("/health-check", handler.Healthcheck)
}
