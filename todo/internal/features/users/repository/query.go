package repository

import (
	"todo/internal/features/users"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(connection *gorm.DB) users.UQuery {
	return &UserQuery{
		db: connection,
	}
}

func (uq *UserQuery) Login(username string) (users.User, error) {
	var result User
	err := uq.db.Where("username = ?", username).First(&result).Error
	if err != nil {
		return users.User{}, err
	}
	return ToUserEntity(result), nil
}

func (uq *UserQuery) Register(newUser users.User) error {
	cnv := ToUserRepo(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		return err
	}
	return nil
}

func (uq *UserQuery) GetUser(userID uint) (users.User, error) {
	var result User
	err := uq.db.Where("id = ?", userID).First(&result).Error
	if err != nil {
		return users.User{}, err
	}
	return ToUserEntity(result), nil
}

func (uq *UserQuery) UpdateUser(updatedUser users.User, userID uint) error {
	cnv := ToUserRepo(updatedUser)
	err := uq.db.Where("id = ?", userID).Updates(&cnv).Error
	if err != nil {
		return err
	}
	return nil
}

func (uq *UserQuery) DeleteUser(userID uint) error {
	var result User
	err := uq.db.Where("id = ?", userID).Delete(&result).Error
	if err != nil {
		return err
	}
	return nil
}
