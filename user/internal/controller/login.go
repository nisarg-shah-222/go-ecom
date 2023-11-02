package controller

import (
	"errors"
	"time"
	"user/internal/constants"
	"user/internal/models"
	"user/internal/models/dtos"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func Login(env *dtos.Env, userLoginRequestDTOs dtos.UserLoginRequestDTOs) (string, error) {
	filter := map[string]interface{}{
		"username": *userLoginRequestDTOs.Username,
	}
	existingUserList := []models.User{}
	env.MySQLConn.Table(constants.TableNameUser).Where(filter).Find(&existingUserList)
	if len(existingUserList) == 0 {
		return "", errors.New("username does not exists")
	}
	if !comparePasswords(existingUserList[0].PasswordHash, []byte(*userLoginRequestDTOs.Password)) {
		return "", errors.New("incorrect password")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["go-ecom"] = existingUserList[0]
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, _ := token.SignedString([]byte(viper.GetString(constants.JWT_SECRET)))
	return tokenString, nil
}
