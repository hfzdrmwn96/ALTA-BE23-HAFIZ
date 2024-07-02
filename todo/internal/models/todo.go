package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Activity string `json:"activity"`
	UserID   uint   `gorm:"primaryKey" json:"userID"`
}

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel {
	return &TodoModel{
		db: connection,
	}
}

func (tm *TodoModel) AddTodo(newTodo Todo) (bool, error) {
	err := tm.db.Create(&newTodo).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (tm *TodoModel) GetTodos(userID uint) ([]Todo, error) {
	var result []Todo
	err := tm.db.Where("user_id = ?", userID).Find(&result).Error
	if err != nil {
		return []Todo{}, err
	}
	return result, nil
}

func (tm *TodoModel) GetTodo(ID uint, userID uint) (Todo, error) {
	var result Todo
	err := tm.db.Where("id = ? AND user_id = ?", ID, userID).First(&result).Error
	if err != nil {
		return Todo{}, err
	}
	return result, nil
}

func (tm *TodoModel) UpdateTodo(newTodo Todo) (Todo, error) {

	err := tm.db.Save(&newTodo).Error
	if err != nil {
		return Todo{}, err
	}
	return newTodo, nil
}

func (tm *TodoModel) DeleteTodo(userID uint, id uint) error {
	err := tm.db.Where("user_id = ? AND id = ?", userID, id).Delete(&Todo{}).Error

	if err != nil {
		return err
	}

	return nil
}
