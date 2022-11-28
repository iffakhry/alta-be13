package repository

import (
	"be13/clean/features/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	FullName  string
	Email     string `gorm:"unique"`
	Password  string
	Role      string `gorm:"type:enum('default', 'admin');default:'default'"`
	Status    bool
	TeamID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//DTO

func (data User) toCore() auth.Core {
	return auth.Core{
		ID:        int(data.ID),
		FullName:  data.FullName,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
