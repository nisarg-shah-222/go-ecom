package dtos

type AuthDtos struct {
	Type     string  `json:"type"`
	Id       *uint   `json:"id"`
	Username *string `json:"username"`
	Role     *string `json:"role"`
}
