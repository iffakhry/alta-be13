package controllers

import (
	"be13/mvc/entities"
	"be13/mvc/helper"
	"be13/mvc/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	authInput := entities.AuthRequest{}
	errBind := c.Bind(&authInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	token, err := repositories.Login(authInput.Email, authInput.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed login"))
	}
	data := map[string]any{
		"token": token,
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", data))
}
