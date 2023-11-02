package middleware

import (
	"product/datastore"
	"product/internal/models/dtos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func EnvMiddleware(ctx *gin.Context) {
	mySQLConn := initialiseDatastores()
	validate := validator.New()
	zapLogger, _ := zap.NewProductionConfig().Build()
	env := &dtos.Env{
		Logger:    zapLogger,
		MySQLConn: mySQLConn,
		Validator: validate,
		Ctx:       ctx,
	}
	ctx.Set("env", env)
	ctx.Next()
}

func initialiseDatastores() *gorm.DB {
	datastore.Get()
	return datastore.MySQLConn
}
