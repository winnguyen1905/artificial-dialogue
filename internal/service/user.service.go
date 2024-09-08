package service

import "artificial-dialogue/internal/repository"

type UserService interface {
	GetUserById()
}

type userService struct {
	userRepository repository.UserRepository
}

func (us *userService) GetUserById() {
	// Implementation for fetching user data
}

func NewUserService() UserService {
	return &userService{
        userRepository: repository.NewUserRepository(),
    }
}