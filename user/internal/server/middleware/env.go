package middleware

import (
	"user/datastore"
	"user/internal/models/dtos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func EnvMiddleware(ctx *gin.Context) {
	mySQLConn := initialiseDatastores()
	validate := validator.New()
	zapLogger, _ := zap.NewProductionConfig().Build()
	decoder := schema.NewDecoder()
	env := &dtos.Env{
		Logger:    zapLogger,
		MySQLConn: mySQLConn,
		Validator: validate,
		Ctx:       ctx,
		Decoder:   decoder,
	}
	ctx.Set("env", env)
	ctx.Next()
}

func initialiseDatastores() *gorm.DB {
	datastore.Get()
	return datastore.MySQLConn
}
