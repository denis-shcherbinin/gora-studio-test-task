package http

import (
	v1 "github.com/denis-shcherbinin/gora-studio-test-task/internal/delivery/http/v1"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"net/http"

	_ "github.com/denis-shcherbinin/gora-studio-test-task/docs"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Static("/stat-img", "./image")

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("api")
	{
		handlerV1.Init(api)
	}
}
