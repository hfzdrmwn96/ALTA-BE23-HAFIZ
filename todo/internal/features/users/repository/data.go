package repository

import (
	t_rep "todo/internal/features/todos/repository"
	"todo/internal/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Email    string
	Phone    string
	Todos    []t_rep.Todo `gorm:"foreignKey:UserID"`
}

func ToUserRepo(input users.User) User {
	return User{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
	}
}

func ToUserEntity(input User) users.User {
	return users.User{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
	}
}
