package handler

import (
	"strconv"
	"todo/internal/features/todos"
	"todo/internal/helper"
	"todo/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodosHand struct {
	srv todos.TService
	tu  utils.TokenUtilInterface
}

func NewTodoHand(s todos.TService, t utils.TokenUtilInterface) todos.THandler {
	return &TodosHand{
		srv: s,
		tu:  t,
	}
}

func (th *TodosHand) AddTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodeToken(c.Get("user").(*jwt.Token))
		var input AddTodoRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil))
		}
		input2 := AddToTodo(input)
		input2.UserID = userID
		err = th.srv.AddTodo(input2)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
	}
}

func (th *TodosHand) GetTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodeToken(c.Get("user").(*jwt.Token))
		todoID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil))
		}

		result, err := th.srv.GetTodo(uint(todoID), userID)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success get users todo data", result))
	}
}

func (th *TodosHand) GetAllTodos() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodeToken(c.Get("user").(*jwt.Token))

		result, err := th.srv.GetAllTodos(userID)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success get all users todos data", result))
	}
}

func (th *TodosHand) UpdateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodeToken(c.Get("user").(*jwt.Token))
		todoID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil))
		}
		var input AddTodoRequest
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil))
		}
		err = th.srv.UpdateTodo(AddToTodo(input), uint(todoID), uint(userID))
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success update data", nil))
	}
}

func (th *TodosHand) DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodeToken(c.Get("user").(*jwt.Token))
		todoID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil))
		}
		err = th.srv.DeleteTodo(uint(todoID), userID)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success delete users todo data", nil))
	}
}
