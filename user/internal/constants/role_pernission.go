package constants

const (
	GUEST = "Guest"
	ADMIN = "Admin"
)

const (
	PATCH_USER = "PATCH_USER"
)

var RolePermissionsMap = map[string][]string{
	ADMIN: {PATCH_USER},
}

var XApiKeyServiceMap = map[string]string{}
