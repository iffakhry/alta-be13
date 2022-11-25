package models

import (
	"be13/mvc/entities"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string `gorm:"type:varchar(15)"`
	Address  string
	Role     string
	Books    []Book
}

func UserCoreToModel(dataCore entities.UserCore) User {
	userGorm := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
		Role:     dataCore.Role,
	}
	return userGorm
}

func ModelToUserCore(dataModel User) entities.UserCore {
	return entities.UserCore{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Phone:     dataModel.Phone,
		Address:   dataModel.Address,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func ListModelToUserCore(dataModel []User) []entities.UserCore {
	var dataCore []entities.UserCore
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToUserCore(v))
	}
	return dataCore
}
