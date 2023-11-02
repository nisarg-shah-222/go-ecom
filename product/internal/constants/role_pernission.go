package constants

const (
	CLIENT = "CLIENT"
	ADMIN  = "ADMIN"
	ORDER  = "ORDER"
)

const (
	GET_PRODUCT    = "GET_PRODUCT"
	UPDATE_PRODUCT = "UPDATE_PRODUCT"
	ADD_PRODUCT    = "ADD_PRODUCT"
)

var RolePermissionsMap = map[string][]string{
	CLIENT: {GET_PRODUCT},
	ADMIN:  {GET_PRODUCT, ADD_PRODUCT, UPDATE_PRODUCT},
	ORDER:  {GET_PRODUCT},
}

var XApiKeyServiceMap = map[string]string{
	"018b399d-fd63-76db-a858-e4b67141a81d": ORDER,
}
