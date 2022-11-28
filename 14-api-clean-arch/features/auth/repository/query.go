package repository

import (
	"be13/clean/features/auth"
	"be13/clean/middlewares"
	"errors"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.RepositoryInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) FindUser(email, password string) (token string, err error) {
	var userData User
	tx := repo.db.Where("email = ? AND password = ?", email, password).First(&userData)
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
