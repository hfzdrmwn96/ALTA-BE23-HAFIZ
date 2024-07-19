package handler

import (
	"todo/internal/features/users"
	"todo/internal/helper"
	"todo/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UsersHand struct {
	srv users.UService
	tu  utils.TokenUtilInterface
}

func NewUserHand(s users.UService, t utils.TokenUtilInterface) users.UHandler {
	return &UsersHand{
		srv: s,
		tu:  t,
	}
}

func (uh *UsersHand) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil))
		}
		err = uh.srv.Register(RegisterToUser(input))
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
	}
}

func (uh *UsersHand) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		result, err := uh.srv.Login(input.Username, input.Password)

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success login", ToLoginReponse(result)))
	}
}

func (uh *UsersHand) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uh.tu.DecodeToken(c.Get("user").(*jwt.Token))

		result, err := uh.srv.GetUser(userID)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success get user data", result))
	}
}

func (uh *UsersHand) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uh.tu.DecodeToken(c.Get("user").(*jwt.Token))

		var input RegisterRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil))
		}
		err = uh.srv.UpdateUser(RegisterToUser(input), uint(userID))
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success update data", nil))
	}
}

func (uh *UsersHand) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uh.tu.DecodeToken(c.Get("user").(*jwt.Token))
		err := uh.srv.DeleteUser(userID)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success delete user data", nil))
	}
}
