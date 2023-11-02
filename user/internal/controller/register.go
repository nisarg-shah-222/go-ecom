package controller

import (
	"errors"
	"user/internal/constants"
	"user/internal/models"
	"user/internal/models/dtos"
)

func Register(env *dtos.Env, userRegisterRequestDTOs dtos.UserRegisterRequestDTOs) error {
	filter := map[string]interface{}{
		"username": *userRegisterRequestDTOs.Username,
	}
	existingUserList := []models.User{}
	env.MySQLConn.Table(constants.TableNameUser).Where(filter).Find(&existingUserList)
	if len(existingUserList) > 0 {
		return errors.New("username already exists")
	}
	hashedPassword := hashAndSalt([]byte(*userRegisterRequestDTOs.Password))
	newUser := models.User{
		Username:     *userRegisterRequestDTOs.Username,
		PasswordHash: hashedPassword,
		Role:         "CLIENT",
	}
	result := env.MySQLConn.Table(constants.TableNameUser).Create(&newUser)
	if result.Error != nil {
		return errors.Join(errors.New("unable to register user"), result.Error)
	}
	return nil
}
