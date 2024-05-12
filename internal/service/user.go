package service

import (
	"context"

	"github.com/mestvv/NNBlogBackend/internal/repository"
)

type userService struct {
	userRepository repository.Users
}

func newUserService(userRepository repository.Users) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

type UserCreateInput struct {
	Name  string
	Email string
}

func (s *userService) Create(ctx context.Context, input UserCreateInput) error {
	return nil
}
