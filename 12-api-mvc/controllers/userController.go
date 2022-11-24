package controllers

import (
	"be13/mvc/entities"
	"be13/mvc/helper"
	"be13/mvc/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {
	result, err := repositories.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	var dataResponse = entities.ListUserCoreToResponse(result)
	// var dataResponse []entities.UserResponse
	// for _, v := range result {
	// 	dataResponse = append(dataResponse, entities.UserResponse{
	// 		Name:    v.Name,
	// 		Email:   v.Email,
	// 		Phone:   v.Phone,
	// 		Address: v.Address,
	// 	})
	// }

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", dataResponse))
}

func AddUserController(c echo.Context) error {
	userInput := entities.UserRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	//mapping
	userCoreData := entities.UserCore{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
		Phone:    userInput.Phone,
		Address:  userInput.Address,
	}

	errInsert := repositories.InsertUser(userCoreData)
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed insert data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success create user"))
}
