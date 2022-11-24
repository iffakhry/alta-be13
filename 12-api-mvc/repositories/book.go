package repositories

import (
	"be13/mvc/config"
	"be13/mvc/entities"
	"be13/mvc/models"
	"errors"
)

func InsertBook(dataCore entities.BookCore) error {
	// userGorm := models.User{
	// 	Name:     dataCore.Name,
	// 	Email:    dataCore.Email,
	// 	Password: dataCore.Password,
	// 	Phone:    dataCore.Phone,
	// 	Address:  dataCore.Address,
	// }
	bookGorm := models.BookCoreToModel(dataCore) //mapping dari core ke struct gorm model, untuk selanjutnya diproses ke query create/insert
	tx := config.DB.Create(&bookGorm)            // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("Insert failed")
	}
	return nil
}

func GetAllBook() ([]entities.BookCore, error) {
	var books []models.Book

	tx := config.DB.Preload("User").Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = models.ListModelToBookCore(books)

	return dataCore, nil
}
