package auth

import "time"

type Core struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	Login(email, password string) (token string, err error)
}

type RepositoryInterface interface {
	FindUser(email, password string) (token string, err error)
}
