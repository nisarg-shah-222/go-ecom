package constants

const (
	GUEST = "Guest"
	ADMIN = "Admin"
)

const (
	CREATE_ORDER = "CREATE_ORDER"
	GET_ORDER    = "GET_ORDER"
)

var RolePermissionsMap = map[string][]string{
	GUEST: {CREATE_ORDER, GET_ORDER},
	ADMIN: {GET_ORDER},
}

var XApiKeyServiceMap = map[string]string{}
