package repositories

import (
	"be13/mvc/config"
	"be13/mvc/entities"
	"be13/mvc/models"
	"errors"
)

func GetAllUser() ([]entities.UserCore, error) {
	var users []models.User

	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore []entities.UserCore
	for _, v := range users {
		dataCore = append(dataCore, entities.UserCore{
			ID:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Password:  v.Password,
			Phone:     v.Phone,
			Address:   v.Address,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})

	}

	return dataCore, nil
}

func InsertUser(dataCore entities.UserCore) error {
	// userGorm := models.User{
	// 	Name:     dataCore.Name,
	// 	Email:    dataCore.Email,
	// 	Password: dataCore.Password,
	// 	Phone:    dataCore.Phone,
	// 	Address:  dataCore.Address,
	// }
	userGorm := models.UserCoreToModel(dataCore)
	tx := config.DB.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("Insert failed")
	}
	return nil
}
