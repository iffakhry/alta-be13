package routes

import (
	"be13/mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetAllUserController)
	e.POST("/users", controllers.AddUserController)
	// e.GET("/users/:id/books", controllers.AddUserController)

	e.GET("/books", controllers.GetAllBookController)
	e.POST("/books", controllers.AddBookController)
	// e.GET("/books/:id", controllers.AddUserController)
	// e.PUT("/books/:id", controllers.AddBookController)
	// e.DELETE("/books/:id", controllers.DeleteBookController)

	return e
}
