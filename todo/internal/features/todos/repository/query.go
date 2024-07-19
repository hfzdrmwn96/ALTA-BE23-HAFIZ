package repository

import (
	"todo/internal/features/todos"

	"gorm.io/gorm"
)

type TodoQuery struct {
	db *gorm.DB
}

func NewTodoQuery(connection *gorm.DB) todos.TQuery {
	return &TodoQuery{
		db: connection,
	}
}

func (tq *TodoQuery) AddTodo(newTodo todos.Todo) error {
	cnv := ToTodoRepo(newTodo)
	err := tq.db.Create(&cnv).Error
	if err != nil {
		return err
	}
	return nil
}

func (tq *TodoQuery) GetAllTodos(userID uint) ([]todos.Todo, error) {
	var result []Todo
	var result2 []todos.Todo
	err := tq.db.Where("user_id = ?", userID).Find(&result).Error
	if err != nil {
		return []todos.Todo{}, err
	}

	for _, v := range result {
		result2 = append(result2, ToTodoEntity(v))
	}
	return result2, nil
}

func (tq *TodoQuery) GetTodo(todoID uint, userID uint) (todos.Todo, error) {
	var result todos.Todo
	err := tq.db.Where("id = ? AND user_id = ?", todoID, userID).First(&result).Error
	if err != nil {
		return todos.Todo{}, err
	}
	return result, nil
}

func (tq *TodoQuery) UpdateTodo(updatedTodo todos.Todo, todoID uint, userID uint) error {
	cnv := ToTodoRepo(updatedTodo)
	err := tq.db.Where("id = ? AND user_id = ?", todoID, userID).Updates(&cnv).Error
	if err != nil {
		return err
	}
	return nil
}

func (tq *TodoQuery) DeleteTodo(todoID uint, userID uint) error {
	err := tq.db.Where("user_id = ? AND id = ?", userID, todoID).Delete(&Todo{}).Error

	if err != nil {
		return err
	}

	return nil
}
