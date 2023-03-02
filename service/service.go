package service

import (
	"context"
	"errors"
	"mockingtest/domain"
	"mockingtest/model"
	repo "mockingtest/repo"
)

var (
	ErrInvalidName     = errors.New("invalid name")
	ErrInvalidEmail    = errors.New("invalid email address")
	ErrInvalidPassword = errors.New("invalid password")
)

type UserService interface {
	AddUser(user domain.AddUserRequest) (userId string, err error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) userService {
	return userService{
		userRepo: userRepo,
	}
}

func (s *userService) AddUser(user domain.AddUserRequest) (userId string, err error) {

	if len(user.Name) < 1 {
		return "", ErrInvalidName
	}

	if len(user.Email) < 1 {
		return "", ErrInvalidEmail
	}

	if len(user.Password) < 6 {
		return "", ErrInvalidPassword
	}

	userModel := model.NewUser(user.Name, user.Email, user.Password)

	userId, err = s.userRepo.AddUser(context.Background(), userModel)
	return
}
