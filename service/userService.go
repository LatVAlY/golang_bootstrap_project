package service

import (
	"golang_bootstrap_project/model"
	"golang_bootstrap_project/repository"
)

type UserService interface {
	FindActiveUserWithUserId(userId string) (*model.UserDTO, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepo
}

func (u UserServiceImpl) FindActiveUserWithUserId(userId string) (*model.UserDTO, error) {
	return u.UserRepository.FindActiveUserWithUserId(userId)
}
