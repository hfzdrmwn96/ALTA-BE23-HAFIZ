package todos

import "todo/internal/models"

type AddTodoRequest struct {
	Activity string `json:"activity"`
	UserID   uint   `json:"userID"`
}

type GetTodoRequest struct {
	UserID uint `json:"userID"`
}

func ToModelTodos(r AddTodoRequest) models.Todo {
	return models.Todo{
		Activity: r.Activity,
		UserID:   r.UserID,
	}
}
