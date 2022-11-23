package repositories

import (
	"be13/mvc/config"
	"be13/mvc/models"
	"errors"
)

func GetAllUser() ([]models.User, error) {
	var users []models.User

	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func InsertUser(user models.User) error {
	tx := config.DB.Create(&user) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("Insert failed")
	}
	return nil
}
