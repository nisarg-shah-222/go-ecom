package constants

const (
	CLIENT = "CLIENT"
	ORDER  = "ORDER"
)

const (
	CREATE_ORDER = "CREATE_ORDER"
)

var RolePermissionsMap = map[string][]string{
	CLIENT: {CREATE_ORDER},
}

var XApiKeyServiceMap = map[string]string{
	"cd67bfbf-28d1-43ee-80ce-f2b1aa97ae0b": ORDER,
}
