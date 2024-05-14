package service

import (
	"context"
	"log/slog"

	"github.com/mestvv/NNBlogBackend/internal/config"
	"github.com/mestvv/NNBlogBackend/internal/repository"
	"github.com/mestvv/NNBlogBackend/pkg/auth"
	"github.com/mestvv/NNBlogBackend/pkg/email"
	"github.com/mestvv/NNBlogBackend/pkg/hash"
	"github.com/mestvv/NNBlogBackend/pkg/otp"
)

type Services struct {
	Users Users
}

type Deps struct {
	Logger       *slog.Logger
	Config       *config.Config
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
	OtpGenerator otp.Generator
	EmailSender  email.Sender
	Repos        *repository.Repositories
}

func NewServices(deps Deps) *Services {
	emailService := newEmailsService(deps.EmailSender, deps.Config.Email)

	return &Services{
		Users: newUserService(deps.Repos.Users,
			deps.Hasher,
			deps.TokenManager,
			deps.OtpGenerator,
			emailService,
			deps.Config.Auth,
		),
	}
}

type Emails interface {
	SendUserVerificationEmail(VerificationEmailInput) error
}

type Users interface {
	Register(ctx context.Context, input UserRegisterInput) error
}
