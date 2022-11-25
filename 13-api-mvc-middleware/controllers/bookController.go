package controllers

import (
	"be13/mvc/entities"
	"be13/mvc/helper"
	"be13/mvc/middlewares"
	"be13/mvc/repositories"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddBookController(c echo.Context) error {
	bookInput := entities.BookRequest{}
	errBind := c.Bind(&bookInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	var dataCore = entities.BookRequestToCore(bookInput) //mapping dari req ke core, untuk selanjutnya dikirim kan ke repository
	errInsert := repositories.InsertBook(dataCore)
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed insert data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success create book"))
}

func GetAllBookController(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idtoken", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roletoken", roleToken)
	if roleToken != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("data hanya bisa dibaca oleh admin"))
	}

	result, err := repositories.GetAllBook()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all books", result))
}
