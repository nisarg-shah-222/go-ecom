package dtos

type AuthDtos struct {
	Id       *uint   `json:"id"`
	Username *string `json:"username"`
	Role     *string `json:"role"`
}
