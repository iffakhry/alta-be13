package repository

import (
	"be13/clean/features/user"
	"database/sql"
)

type userRawRepository struct {
	db *sql.DB
}

func NewRaw(db *sql.DB) user.RepositoryInterface {
	return &userRawRepository{
		db: db,
	}
}

// GetById implements user.RepositoryInterface
func (*userRawRepository) GetById(id int) (data user.Core, err error) {
	panic("unimplemented")
}

// Create implements user.RepositoryInterface
func (*userRawRepository) Create(input user.Core) (row int, err error) {
	panic("unimplemented")
}

// GetAll implements user.RepositoryInterface
func (*userRawRepository) GetAll() (data []user.Core, err error) {
	panic("unimplemented")
}
