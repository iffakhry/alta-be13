package repository

import (
	_user "be13/clean/features/user"

	"gorm.io/gorm"
)

//struct gorm model
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string `gorm:"type:varchar(15)"`
	Address  string
	Role     string
	// Books    []Book
}

// DTO
// mapping

//mengubah struct core ke struct model gorm
func fromCore(dataCore _user.Core) User {
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

//mengubah struct model gorm ke struct core
func (dataModel *User) toCore() _user.Core {
	return _user.Core{
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

//mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []User) []_user.Core {
	var dataCore []_user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
