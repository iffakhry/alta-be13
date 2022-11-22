package main

// go get -u github.com/labstack/echo/v4

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// ID      int    `json:"id" form:"id"`
	// Title   string `json:"title" form:"title"`
	// Content string `json:"content" form:"content"`
}

type User struct {
	ID          int    `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	// ID          int    `json:"id"`
	// Name        string `json:"name"`
	// Email       string `json:"email"`
	// Password    string `json:"password"`
	// PhoneNumber string `json:"phone_number"`
}

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/hello", HelloController)
	e.POST("/hello", PostHelloController)
	e.GET("/articles", GetArticlesController)
	e.GET("/articles/:id", GetArticlesByIdController)
	e.POST("/articles", AddArticlesController)
	e.POST("/users", AddUsersController)
	e.PUT("/users/:id", UpdateUsersController)
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

//contoh query param
func GetArticlesController(c echo.Context) error {
	pageQuery := c.QueryParam("page")
	titleQuery := c.QueryParam("title")
	// if pageQuery == "" {
	// 	pageQuery = "1"
	// }
	// limit := c.QueryParam("limit")
	var data = []Article{
		{1, "Title 1", "Content 1"},
		{2, "Title 2", "Content 2"},
		{3, "Title 3", "Content 3"},
		{4, "Title 4", "Content 4"},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success read data",
		"title":   titleQuery,
		"page":    pageQuery,
		"data":    data,
	})
}

//contoh param
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

func AddArticlesController(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	var data = Article{
		ID:      1,
		Title:   title,
		Content: content,
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"data":    data,
	})
}

// memanfaatkan bind
func AddUsersController(c echo.Context) error {
	// var user User
	user := User{}
	errBind := c.Bind(&user)
	// c.Request().FormFile("foto")
	fmt.Println("error", errBind)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error binding data" + errBind.Error(),
		})
	}

	if user.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "nama harus diisi",
		})
	}

	fmt.Println("nama user", user.Name)
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"data":    user,
	})
}

func UpdateUsersController(c echo.Context) error {
	// var user User
	idParam := c.Param("id") //id yang akan diupdate
	id, _ := strconv.Atoi(idParam)

	user := User{}
	errBind := c.Bind(&user) //data yang diupdate
	// c.Request().FormFile("foto")
	fmt.Println("error", errBind)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error binding data" + errBind.Error(),
		})
	}

	if user.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "nama harus diisi",
		})
	}

	fmt.Println("nama user", user.Name)
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"data":    user,
		"id":      id,
	})
}
