package api

import (
	"log/slog"

	sloggin "github.com/samber/slog-gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/mestvv/NNBlogBackend/docs"

	"github.com/gin-gonic/gin"
	internalV1 "github.com/mestvv/NNBlogBackend/internal/api/internal/http/v1"
	"github.com/mestvv/NNBlogBackend/internal/config"
	"github.com/mestvv/NNBlogBackend/internal/service"
)

type Handler struct {
	services *service.Services
	logger   *slog.Logger
}

func NewHandlers(services *service.Services, logger *slog.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(
		gin.Recovery(),
		sloggin.NewWithConfig(h.logger, sloggin.Config{
			WithSpanID:  true,
			WithTraceID: true,
		}),
	)

	if cfg.HttpServer.SwaggerEnabled {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("internal")))
	}

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	internalHandlersV1 := internalV1.NewHandler(h.services, h.logger)
	api := router.Group("/api")
	{
		internalHandlersV1.Init(api)
	}
}
