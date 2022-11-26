package service

import "be13/clean/features/user"

type userService struct {
	userRepository user.RepositoryInterface
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
	}
}

// Create implements user.ServiceInterface
func (service *userService) Create(input user.Core) (err error) {
	panic("unimplemented")
}

// GetAll implements user.ServiceInterface
func (service *userService) GetAll() (data []user.Core, err error) {
	data, err = service.userRepository.GetAll()
	return

}
