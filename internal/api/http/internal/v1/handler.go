package v1

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/mestvv/NNBlogBackend/internal/service"
	"github.com/mestvv/NNBlogBackend/pkg/auth"
)

// @title Backend API
// @version 1.0
// @description Backend API for NNBlog Service

// @BasePath /api/v1

// @securityDefinitions.apikey AdminAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey UserAuth
// @in header
// @name Authorization

type Handler struct {
	services     *service.Services
	logger       *slog.Logger
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, logger *slog.Logger, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		logger:       logger,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("v1")
	{
		h.initUsersRoutes(v1)
	}
}
