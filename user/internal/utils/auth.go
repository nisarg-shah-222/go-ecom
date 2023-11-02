package utils

import "user/internal/models/dtos"

func HasPermission(env *dtos.Env, permission string) bool {
	permissionsList := env.PermissionList
	for _, val := range permissionsList {
		if val == permission {
			return true
		}
	}
	return false
}
