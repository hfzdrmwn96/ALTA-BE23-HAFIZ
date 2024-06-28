package users

import "todo/internal/models"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string          `json:"name"`
	Password string          `json:"password"`
	Email    string          `json:"email"`
	Phone    string          `json:"hp"`
	Address  []AlamatRequest `json:"address"`
}

type AlamatRequest struct {
	Alamat string `json:"alamat"`
}

func ToModelUsers(r RegisterRequest) models.User {
	return models.User{
		Name:     r.Name,
		Password: r.Password,
		Email:    r.Email,
		Phone:    r.Phone,
	}
}
