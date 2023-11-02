package controller

import (
	"errors"
	"user/internal/constants"
	"user/internal/models/dtos"
)

func UpdateUser(env *dtos.Env, userRegisterRequestDTOs dtos.UpdateUserRequestDTOs) error {
	filter := map[string]interface{}{
		"id": *userRegisterRequestDTOs.Id,
	}
	updates := map[string]interface{}{
		"role": *userRegisterRequestDTOs.Role,
	}
	result := env.MySQLConn.Table(constants.TableNameUser).Where(filter).Updates(updates)
	if result.Error != nil {
		return errors.Join(errors.New("unable to update user"), result.Error)
	}
	return nil
}
