package repository

import (
	"be13/clean/features/user"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *userRepository) Create(input user.Core) (row int, err error) {
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *userRepository) GetAll() (data []user.Core, err error) {
	var users []User

	tx := repo.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(users)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (*userRepository) GetById(id int) (data user.Core, err error) {
	panic("unimplemented")
}
