package routes

import (
	"todo/configs"
	"todo/internal/features/todos"
	"todo/internal/features/users"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.UHandler, th todos.THandler) {
	e.POST("/login", uh.Login())
	e.POST("/register", uh.Register())
	UsersRoute(e, uh)
	TodosRoute(e, th)
}

func UsersRoute(e *echo.Echo, uh users.UHandler) {
	u := e.Group("/users")
	u.Use(JWTConfig())
	u.GET("", uh.GetUser())
	u.PUT("", uh.UpdateUser())
}

func TodosRoute(e *echo.Echo, th todos.THandler) {
	t := e.Group("/todos")
	t.Use(JWTConfig())
	t.GET("", th.GetAllTodos())
	t.GET("/:id", th.GetTodo())
	t.POST("", th.AddTodo())
	t.PUT("/:id", th.UpdateTodo())
	t.DELETE("/:id", th.DeleteTodo())

}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(configs.ImportSetting().Passkey),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
