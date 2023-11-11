package dtos

type UserRegisterRequestDTOs struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required,gt=7"`
}

type UserLoginRequestDTOs struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required,gt=7"`
}

type UpdateUserRequestDTOs struct {
	Id   *uint   `json:"id" validate:"required"`
	Role *string `json:"role" validate:"required,oneof=Guest Admin"`
}
