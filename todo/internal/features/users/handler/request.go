package handler

import "todo/internal/features/users"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func LoginToUser(input LoginRequest) users.User {
	return users.User{
		Username: input.Username,
		Password: input.Password,
	}
}

func RegisterToUser(input RegisterRequest) users.User {
	return users.User{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
	}
}
