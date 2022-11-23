package routes

import (
	"be13/mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetAllUserController)
	e.POST("/users", controllers.AddUserController)

	return e
}
