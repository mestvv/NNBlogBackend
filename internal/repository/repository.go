package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mestvv/NNBlogBackend/internal/domain"
)

type Repositories struct {
	Users Users
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: newUserRepository(db),
	}
}

type Users interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, ID uuid.UUID) (*domain.User, error)
}
