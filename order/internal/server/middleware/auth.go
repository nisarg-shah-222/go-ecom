package middleware

import (
	"errors"
	"net/http"
	"order/internal/constants"
	"order/internal/models/dtos"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func AuthMiddleware(ctx *gin.Context) {
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
	env.Ctx = ctx
	ctx.Status(http.StatusInternalServerError)
	requestAuthToken := ctx.GetHeader("Authorization")
	xApiKey := ctx.GetHeader("X-Api-Key")
	if requestAuthToken == "" && xApiKey == "" {
		env.Logger.Error("authorization failed, auth token and x-api-key both not present")
		AbortWith401(ctx)
		return
	}
	if xApiKey != "" {
		permissionList, err := getPermissionListFromXApiKey(env, &xApiKey)
		if err != nil {
			env.Logger.Error("authorization failed")
			AbortWith401(ctx)
			return
		}
		env.PermissionList = permissionList
		env.Logger.Info("Request Authenticated")
	} else {
		token, authDtos, err := validateAuthToken(requestAuthToken)
		permissionList := constants.RolePermissionsMap[*authDtos.Role]
		env.PermissionList = permissionList
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 1, "message": "error", "error": errors.Join(err, errors.New("unauthorized access")).Error()})
			return
		}
		if !token.Valid {
			env.Logger.Warn("Invalid token received", zap.Any("token", requestAuthToken))
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 1, "message": "error", "error": "token invalid"})
			return
		}
		env.Logger.Info("Request Authenticated", zap.Any("username", authDtos.Username))
		env.AuthDtos = authDtos
	}
	ctx.Next()
}

func getPermissionListFromXApiKey(env *dtos.Env, xApiKey *string) ([]string, error) {
	if xApiKey == nil || *xApiKey == "" {
		env.Logger.Error("x-api-key not present")
		return nil, errors.New("x-api-key not present")
	}
	serviceKeyMap := constants.XApiKeyServiceMap
	if service, ok := serviceKeyMap[*xApiKey]; ok {
		return constants.RolePermissionsMap[service], nil
	}
	env.Logger.Error("x-api-key did not matched")
	return nil, errors.New("request denied to the service")
}

func validateAuthToken(encodedToken string) (*jwt.Token, *dtos.AuthDtos, error) {
	keyFunc := func(_ *jwt.Token) (interface{}, error) {
		jwtSecret := viper.GetString(constants.JWT_SECRET)

		return []byte(jwtSecret), nil
	}

	token, err := jwt.Parse(encodedToken, keyFunc)
	if err != nil {
		return nil, nil, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	userDetailsMap := claims["go-ecom"].(map[string]interface{})
	var id uint
	if authId, ok := userDetailsMap["id"].(float64); ok {
		id = uint(authId)
	}
	var username string = userDetailsMap["username"].(string)
	var role string = userDetailsMap["role"].(string)
	authDtos := &dtos.AuthDtos{
		Id:       &id,
		Username: &username,
		Role:     &role,
	}
	return token, authDtos, nil
}

func AbortWith401(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"status": 1, "message": "error", "error": "token invalid"})
	ctx.AbortWithStatus(http.StatusUnauthorized)
}
