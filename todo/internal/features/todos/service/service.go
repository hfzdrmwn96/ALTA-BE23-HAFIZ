package service

import (
	"todo/internal/features/todos"
)

type TodoSrv struct {
	qry todos.TQuery
}

func NewTodoSrv(q todos.TQuery) todos.TService {
	return &TodoSrv{
		qry: q,
	}
}

func (us *TodoSrv) AddTodo(input todos.Todo) error {
	err := us.qry.AddTodo(input)
	if err != nil {
		return err
	}
	return nil
}

func (us *TodoSrv) GetTodo(todoID uint, userID uint) (todos.Todo, error) {
	result, err := us.qry.GetTodo(todoID, userID)
	if err != nil {
		return todos.Todo{}, err
	}
	return result, nil
}

func (us *TodoSrv) GetAllTodos(userID uint) ([]todos.Todo, error) {

	result, err := us.qry.GetAllTodos(userID)
	if err != nil {
		return []todos.Todo{}, err
	}

	return result, nil
}

func (us *TodoSrv) UpdateTodo(input todos.Todo, todoID uint, userID uint) error {

	err := us.qry.UpdateTodo(input, todoID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (us *TodoSrv) DeleteTodo(todoID uint, userID uint) error {
	err := us.qry.DeleteTodo(todoID, userID)
	if err != nil {
		return err
	}

	return nil
}
