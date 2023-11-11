package server

import (
	"fmt"
	"product/internal/constants"
	"product/internal/handler"
	"product/internal/server/middleware"

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
		privateV1Apis.POST("/product", handler.AddProduct)
		privateV1Apis.POST("/permission", handler.AddProductUserPermission)
		privateV1Apis.GET("/product/:id", handler.GetProduct)
		privateV1Apis.PATCH("/product/:id", handler.UpdateProduct)
		privateV1Apis.PATCH("/product/:id/details", handler.UpdateProductDetails)
		privateV1Apis.DELETE("/product/:id", handler.DeleteProduct)
	}
	public.GET("/health-check", handler.Healthcheck)
}
