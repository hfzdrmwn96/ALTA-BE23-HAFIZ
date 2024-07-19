package factory

import (
	"fmt"
	"todo/configs"
	t_hnd "todo/internal/features/todos/handler"
	t_rep "todo/internal/features/todos/repository"
	t_srv "todo/internal/features/todos/service"

	u_hnd "todo/internal/features/users/handler"
	u_rep "todo/internal/features/users/repository"
	u_srv "todo/internal/features/users/service"
	"todo/internal/routes"
	"todo/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo) {
	db, _ := configs.ConnectDB(configs.ImportSetting())

	MigrateDB(db)

	pu := utils.NewPassUtil()
	tu := utils.NewTokenUtil()

	uq := u_rep.NewUserQuery(db)
	us := u_srv.NewUserSrv(uq, pu, tu)
	uh := u_hnd.NewUserHand(us, tu)

	tq := t_rep.NewTodoQuery(db)
	ts := t_srv.NewTodoSrv(tq)
	th := t_hnd.NewTodoHand(ts, tu)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, uh, th)
}

func MigrateDB(db *gorm.DB) {

	err := db.AutoMigrate(&u_rep.User{}, &t_rep.Todo{})
	if err != nil {
		fmt.Println(err)
	}

}
