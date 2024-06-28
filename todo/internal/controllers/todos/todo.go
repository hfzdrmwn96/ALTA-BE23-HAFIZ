package todos

import (
	"strconv"
	"todo/internal/helper"
	"todo/internal/models"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) AddTodo(c echo.Context) error {
	var input AddTodoRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	_, err = tc.model.AddTodo(ToModelTodos(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
}

func (tc *TodoController) GetTodos(c echo.Context) error {

	userID, err := strconv.Atoi(c.Param("userID"))

	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "user id Error", nil))
	}

	result, err := tc.model.GetTodos(uint(userID))

	if len(result) == 0 {
		return c.JSON(400, helper.ResponseFormat(400, "data empty", nil))
	}

	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}

	return c.JSON(200, helper.ResponseFormat(200, "success", result))
}

func (tc *TodoController) UpdateTodo(c echo.Context) error {

	UID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "user id Error", nil))
	}

	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "user id Error", nil))
	}

	result, err := tc.model.GetTodo(uint(ID), uint(UID))

	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}

	err = c.Bind(&result.Activity)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}

	result.UserID = uint(UID)
	result.ID = uint(ID)
	result, err = tc.model.UpdateTodo(result)

	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}

	return c.JSON(200, helper.ResponseFormat(200, "success", result))

}
