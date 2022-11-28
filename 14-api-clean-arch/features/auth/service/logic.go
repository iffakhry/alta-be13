package service

import "be13/clean/features/auth"

type authService struct {
	authData auth.RepositoryInterface
}

func New(data auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authData: data,
	}
}

func (service *authService) Login(email, password string) (string, error) {
	token, err := service.authData.FindUser(email, password)
	return token, err
}
