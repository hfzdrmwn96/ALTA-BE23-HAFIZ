package users

import "github.com/labstack/echo/v4"

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
}

type UService interface {
	Login(username, password string) (string, error)
	Register(input User) error
	GetUser(userID uint) (User, error)
	UpdateUser(input User, userID uint) error
	DeleteUser(userID uint) error
}

type UQuery interface {
	Login(email string) (User, error)
	Register(newUser User) error
	GetUser(userID uint) (User, error)
	UpdateUser(updatedUser User, userID uint) error
	DeleteUser(userID uint) error
}
