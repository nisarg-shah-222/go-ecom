package dtos

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Env struct {
	Logger         *zap.Logger
	AuthDtos       *AuthDtos
	MySQLConn      *gorm.DB
	Ctx            *gin.Context
	Validator      *validator.Validate
	Decoder        *schema.Decoder
	PermissionList []string
}
