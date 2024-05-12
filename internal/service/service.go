package service

import (
	"context"
	"log/slog"

	"github.com/mestvv/NNBlogBackend/internal/config"
	"github.com/mestvv/NNBlogBackend/internal/repository"
)

type Services struct {
	Users Users
}

type Deps struct {
	Logger *slog.Logger
	Config *config.Config
	Repos  *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Users: newUserService(deps.Repos.Users),
	}
}

type Users interface {
	Create(ctx context.Context, input UserCreateInput) error
}
