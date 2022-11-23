package controllers

import (
	"be13/mvc/helper"
	"be13/mvc/models"
	"be13/mvc/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {
	result, err := repositories.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", result))
}

func AddUserController(c echo.Context) error {
	userInput := models.User{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	errInsert := repositories.InsertUser(userInput)
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed insert data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success create user"))
}
