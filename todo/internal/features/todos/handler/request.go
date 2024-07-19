package handler

import (
	"todo/internal/features/todos"
)

type AddTodoRequest struct {
	Activity string `json:"activity"`
}

func AddToTodo(input AddTodoRequest) todos.Todo {
	return todos.Todo{
		Activity: input.Activity,
	}
}
