package main

import (
	"todo/configs"
	"todo/internal/controllers/todos"
	"todo/internal/controllers/users"
	"todo/internal/models"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	cfg := configs.ImportSetting()
	db, _ := configs.ConnectDB(cfg)

	db.AutoMigrate(&models.User{}, &models.Todo{})

	um := models.NewUserModel(db)
	uc := users.NewUserController(um)

	tm := models.NewTodoModel(db)
	tc := todos.NewTodoController(tm)

	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)
	e.POST("/todos/:userID", tc.AddTodo)
	e.GET("/todos/:userID", tc.GetTodos)
	e.PUT("/todos/:userID/:ID", tc.UpdateTodo)
	e.DELETE("/todos/:userID/:ID", tc.DeleteTodo)

	e.Start(":8000")
}
