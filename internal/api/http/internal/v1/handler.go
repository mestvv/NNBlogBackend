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
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

type Handler struct {
	Services *service.Services
	Logger   *slog.Logger
}

func NewHandler(services *service.Services, logger *slog.Logger) *Handler {
	return &Handler{
		Services: services,
		Logger:   logger,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("v1")
	{
		h.initUsersRoutes(v1)
	}
}
