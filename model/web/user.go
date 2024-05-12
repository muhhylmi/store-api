package web

import "github.com/muhhylmi/store-api/model/domain"

type UserCreateRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=5" json:"password"`
	Role     string `validate:"required,oneof=CUSTOMER ADMIN" json:"role"`
}

type LoginRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=5" json:"password"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"access_token"`
}

func ToUserRersponse(user domain.Users) UserResponse {
	return UserResponse{
		Id:       user.BaseModel.ID,
		Username: user.Username,
		Role:     user.Role,
	}
}

func ToLoginResponse(user domain.Users, token string) LoginResponse {
	return LoginResponse{
		Id:       user.BaseModel.ID,
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}
}
