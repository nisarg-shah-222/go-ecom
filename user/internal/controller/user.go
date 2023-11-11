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

func Register(env *dtos.Env, userRegisterRequestDTOs dtos.UserRegisterRequestDTOs) error {
	filter := map[string]interface{}{
		"username": *userRegisterRequestDTOs.Username,
	}
	existingUserList := []models.User{}
	env.MySQLConn.Table(constants.TableNameUser).Where(filter).Find(&existingUserList)
	if len(existingUserList) > 0 {
		return errors.Join(constants.ErrBadRequest, constants.ErrUserNameExists)
	}
	hashedPassword := hashAndSalt([]byte(*userRegisterRequestDTOs.Password))
	newUser := models.User{
		Username:     *userRegisterRequestDTOs.Username,
		PasswordHash: hashedPassword,
		Role:         "Guest",
	}
	result := env.MySQLConn.Table(constants.TableNameUser).Create(&newUser)
	if result.Error != nil {
		return errors.Join(constants.ErrBadRequest, constants.ErrUserNameExists)
	}
	return nil
}

func Login(env *dtos.Env, userLoginRequestDTOs dtos.UserLoginRequestDTOs) (string, error) {
	filter := map[string]interface{}{
		"username": *userLoginRequestDTOs.Username,
	}
	existingUserList := []models.User{}
	env.MySQLConn.Table(constants.TableNameUser).Where(filter).Find(&existingUserList)
	if len(existingUserList) == 0 {
		return "", errors.Join(constants.ErrBadRequest, constants.ErrUserNameNotExists)
	}
	if !comparePasswords(existingUserList[0].PasswordHash, []byte(*userLoginRequestDTOs.Password)) {
		return "", errors.Join(constants.ErrBadRequest, constants.ErrIncorrectPassword)
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["go-ecom"] = existingUserList[0]
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, _ := token.SignedString([]byte(viper.GetString(constants.JWT_SECRET)))
	return tokenString, nil
}

func UpdateUser(env *dtos.Env, userRegisterRequestDTOs dtos.UpdateUserRequestDTOs) error {
	filter := map[string]interface{}{
		"id": *userRegisterRequestDTOs.Id,
	}
	updates := map[string]interface{}{
		"role": *userRegisterRequestDTOs.Role,
	}
	result := env.MySQLConn.Table(constants.TableNameUser).Where(filter).Updates(updates)
	if result.Error != nil {
		return errors.Join(constants.ErrInternalServerError, result.Error)
	}
	return nil
}
