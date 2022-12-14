package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", HelloController)
	e.GET("/world", WorldController)

	e.Logger.Fatal(e.Start(":8080"))
}

func HelloController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"message": "hello world",
	})
}
func WorldController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"message": "world hello",
	})
}
