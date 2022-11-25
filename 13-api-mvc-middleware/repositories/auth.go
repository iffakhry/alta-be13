package repositories

import (
	"be13/mvc/config"
	"be13/mvc/middlewares"
	"be13/mvc/models"
	"errors"
)

func Login(email, password string) (string, error) {
	var userData models.User
	tx := config.DB.Where("email = ? AND password = ?", email, password).First(&userData)
	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return "", errors.New("login failed")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID), userData.Role)
	if errToken != nil {
		return "", errToken
	}

	return token, nil

}
