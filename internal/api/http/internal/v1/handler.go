package v1

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/mestvv/NNBlogBackend/internal/service"
)

// @title Backend Service
// @version 1.0
// @description API for Service

// @BasePath /api/v1

// @securityDefinitions.apikey AdminAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization

type Handler struct {
	services *service.Services
	logger   *slog.Logger
}

func NewHandler(services *service.Services, logger *slog.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("v1")
	{
		h.initUsersRoutes(v1)
	}
}
