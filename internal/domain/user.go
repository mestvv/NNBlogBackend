package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRole int

const (
	UserRoleDefault UserRole = 0
	UserRoleAdmin   UserRole = 10
	UserRoleModer   UserRole = 20
)

type User struct {
	ID               uuid.UUID `db:"id"`
	Password         string    `db:"password"`
	Username         string    `db:"username"`
	FirstName        *string   `db:"first_name"`
	LastName         *string   `db:"last_name"`
	Email            string    `db:"email"`
	Bio              *string   `db:"bio"`
	Role             UserRole  `db:"role"`
	Active           bool      `db:"active"`
	VerificationCode string    `db:"verification_code"`
	Verified         bool      `db:"verified"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
