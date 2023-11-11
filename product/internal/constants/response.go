package constants

var (
	RESPONSE_FORBIDDEN = map[string]any{
		"status":  1,
		"message": "forbidden",
	}

	RESPONSE_BAD_REQUEST = map[string]any{
		"status":  1,
		"message": "bad request",
	}

	RESPONSE_INTERNAL_SERVER_ERROR = map[string]any{
		"status":  1,
		"message": "our systems are facing problems. please try again after some time",
	}
)
