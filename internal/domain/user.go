package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	UserRoleDefault = "default"
	UserRoleAdmin   = "administrator"
	UserRoleModer   = "moderator"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Password  string    `db:"password"`
	Username  string    `db:"username"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Phone     string    `db:"phone"`
	Bio       string    `db:"bio"`
	Role      UserRole  `db:"role"`
	Active    bool      `db:"active"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
