package service

import (
	"todo/internal/features/users"
	"todo/internal/utils"
)

type UserSrv struct {
	qry users.UQuery
	pu  utils.PassUtilInterface
	tu  utils.TokenUtilInterface
}

func NewUserSrv(q users.UQuery, p utils.PassUtilInterface, t utils.TokenUtilInterface) users.UService {
	return &UserSrv{
		qry: q,
		pu:  p,
		tu:  t,
	}
}

func (us *UserSrv) Login(username, password string) (string, error) {
	result, err := us.qry.Login(username)
	if err != nil {
		return "", err
	}

	err = us.pu.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := us.tu.GenerateToken(result)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserSrv) Register(input users.User) error {
	hashedPass, err := us.pu.GeneratePassword(input.Password)
	if err != nil {
		return err
	}
	input.Password = string(hashedPass)

	err = us.qry.Register(input)

	return err
}

func (us *UserSrv) GetUser(userID uint) (users.User, error) {
	return us.qry.GetUser(userID)
}

func (us *UserSrv) UpdateUser(input users.User, userID uint) error {
	hashedPass, err := us.pu.GeneratePassword(input.Password)
	if err != nil {
		return err
	}
	input.Password = string(hashedPass)

	return us.qry.UpdateUser(input, userID)
}

func (us *UserSrv) DeleteUser(userID uint) error {
	return us.qry.DeleteUser(userID)
}
