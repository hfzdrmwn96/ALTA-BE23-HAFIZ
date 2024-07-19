package repository

import (
	"todo/internal/features/todos"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Activity string
	UserID   uint
}

func ToTodoRepo(input todos.Todo) Todo {
	return Todo{
		Activity: input.Activity,
		UserID:   input.UserID,
	}
}

func ToTodoEntity(input Todo) todos.Todo {
	return todos.Todo{
		ID:       input.ID,
		Activity: input.Activity,
		UserID:   input.UserID,
	}
}
