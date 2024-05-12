package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mestvv/NNBlogBackend/internal/domain"
)

type userRepository struct {
	db *sqlx.DB
}

func newUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByID(ctx, ID uuid.UUID) (*domain.User, error) {
	return nil, nil
}
