package repository

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mestvv/NNBlogBackend/internal/db"
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

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	const query = `
	INSERT INTO user (id, password, username, first_name, last_name, email, role, active, verification_code, verified) VALUES(uuid_to_bin(?), ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Password, user.Username, user.FirstName, user.LastName, user.Email, user.Role, user.Active, user.VerificationCode, user.Verified)

	if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok && mysqlError.Number == db.DuplicateEntry {
			return domain.ErrDuplicateEntry
		}
		return fmt.Errorf("db insert user: %v", err)
	}

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, ID uuid.UUID) (*domain.User, error) {
	return nil, nil
}
