package constants

const (
	GUEST = "Guest"
	ADMIN = "Admin"
	ORDER = "ORDER"
)

const (
	GET_PRODUCT                 = "GET_PRODUCT"
	UPDATE_PRODUCT              = "UPDATE_PRODUCT"
	ADD_PRODUCT                 = "ADD_PRODUCT"
	ADD_PRODUCT_USER_PERMISSION = "ADD_PRODUCT_USER_PERMISSION"
	DELETE_PRODUCT              = "DELETE_PRODUCT"
)

const (
	VIEW   = "View"
	UPDATE = "Update"
	DELETE = "Delete"
)

var PermissionsHierarchyMap = map[string][]string{
	VIEW:   {VIEW},
	UPDATE: {VIEW, UPDATE},
	DELETE: {VIEW, UPDATE, DELETE},
}

var RolePermissionsMap = map[string][]string{
	GUEST: {GET_PRODUCT},
	ADMIN: {GET_PRODUCT, ADD_PRODUCT, UPDATE_PRODUCT, ADD_PRODUCT_USER_PERMISSION, DELETE_PRODUCT},
	ORDER: {GET_PRODUCT},
}

var XApiKeyServiceMap = map[string]string{
	"018b399d-fd63-76db-a858-e4b67141a81d": ORDER,
}
