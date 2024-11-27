package service

import (
	"context"
	"errors"

	"github.com/TimesCoder/movie-app/internal/entity"
	"github.com/TimesCoder/movie-app/internal/repository"
)

type UserService interface {
	Login(ctx context.Context, username, password string) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Login(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("username atau passowrd salah")
	}

	return user, nil
}
