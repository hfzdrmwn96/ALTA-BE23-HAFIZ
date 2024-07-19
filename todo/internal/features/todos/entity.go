package todos

import "github.com/labstack/echo/v4"

type Todo struct {
	ID       uint
	Activity string
	UserID   uint
}

type THandler interface {
	AddTodo() echo.HandlerFunc
	GetTodo() echo.HandlerFunc
	GetAllTodos() echo.HandlerFunc
	UpdateTodo() echo.HandlerFunc
	DeleteTodo() echo.HandlerFunc
}

type TService interface {
	AddTodo(input Todo) error
	GetTodo(todoID uint, userID uint) (Todo, error)
	GetAllTodos(userID uint) ([]Todo, error)
	UpdateTodo(input Todo, todoID uint, userID uint) error
	DeleteTodo(todoID uint, userID uint) error
}

type TQuery interface {
	AddTodo(newTodo Todo) error
	GetAllTodos(userID uint) ([]Todo, error)
	GetTodo(todoID uint, userID uint) (Todo, error)
	UpdateTodo(updatedTodo Todo, todoID uint, userID uint) error
	DeleteTodo(todoID uint, userID uint) error
}
