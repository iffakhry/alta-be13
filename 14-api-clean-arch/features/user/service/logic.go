package service

import (
	"be13/clean/features/user"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *userService) Create(input user.Core) (err error) {
	//validate
	// if input.Name == "" || input.Email == "" || input.Password == "" {
	// 	return errors.New("Name, email, password harus diisi")
	// }
	input.Role = "user"
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements user.ServiceInterface
func (service *userService) GetAll() (data []user.Core, err error) {
	data, err = service.userRepository.GetAll()
	return

}
