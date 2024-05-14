package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mestvv/NNBlogBackend/internal/config"
	"github.com/mestvv/NNBlogBackend/internal/domain"
	"github.com/mestvv/NNBlogBackend/internal/repository"
	"github.com/mestvv/NNBlogBackend/pkg/auth"
	"github.com/mestvv/NNBlogBackend/pkg/hash"
	"github.com/mestvv/NNBlogBackend/pkg/otp"
)

type userService struct {
	userRepository repository.Users
	hasher         hash.PasswordHasher
	tokenManager   auth.TokenManager
	otpGenerator   otp.Generator
	emailService   Emails
	authConfig     config.AuthConfig
}

func newUserService(userRepository repository.Users,
	hasher hash.PasswordHasher,
	tokenManager auth.TokenManager,
	otpGenerator otp.Generator,
	emailService Emails,
	authConfig config.AuthConfig,
) *userService {
	return &userService{
		userRepository: userRepository,
		hasher:         hasher,
		tokenManager:   tokenManager,
		otpGenerator:   otpGenerator,
		emailService:   emailService,
		authConfig:     authConfig,
	}
}

type UserRegisterInput struct {
	Username  string
	FirstName *string
	LastName  *string
	Email     string
	Password  string
}

func (s *userService) Register(ctx context.Context, input UserRegisterInput) error {
	userUUID, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("generate uuid v7 failed: %v", err)
	}

	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return fmt.Errorf("hasher failed: %w", err)
	}

	verificationCode := s.otpGenerator.RandomSecret(s.authConfig.VerificationCodeLength)

	// TODO: верификация пользователя
	user := domain.User{
		ID:               userUUID,
		Password:         passwordHash,
		Username:         input.Username,
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		Email:            input.Email,
		Role:             domain.UserRoleDefault,
		Active:           true,
		VerificationCode: verificationCode,
		Verified:         true,
	}

	if err := s.userRepository.Create(ctx, &user); err != nil {
		if errors.Is(err, domain.ErrDuplicateEntry) {
			return ErrUserAlreadyExist
		}
		return fmt.Errorf("create user failed: %w", err)
	}

	return nil
}
