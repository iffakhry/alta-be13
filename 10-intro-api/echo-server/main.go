package main

// go get -u github.com/labstack/echo/v4

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/hello", HelloController)
	e.POST("/hello", PostHelloController)
	e.GET("/articles", GetArticlesController)
	e.GET("/articles/:id", GetArticlesByIdController)
	// start the server, and log if it fails
	e.Start(":8080")
}

// handler - Simple handler to make sure environment is setup
func HelloController(c echo.Context) error {
	// return the string "Hello World" as the response body
	// with an http.StatusOK (200) status

	// return c.String(http.StatusOK, "Hello World")
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "Hello world",
	})
}

func PostHelloController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "POST Hello world",
	})
}

func GetArticlesController(c echo.Context) error {
	var data = []Article{
		{1, "Title 1", "Content 1"},
		{2, "Title 2", "Content 2"},
		{3, "Title 3", "Content 3"},
		{4, "Title 4", "Content 4"},
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"data":    data,
	})
}

func GetArticlesByIdController(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "id must integer",
		})
	}
	var data = Article{
		ID:      id,
		Title:   "Title 1",
		Content: "Content 1",
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"data":    data,
	})
}
