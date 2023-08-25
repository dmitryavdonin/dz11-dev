package handler

import (
	"notification/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/notification")
	{
		api.POST("/", h.create)
		api.GET("/:id", h.getByOrderId)
		api.GET("/", h.getAll)
		api.DELETE("/:id", h.deleteByOrderId)
	}

	return router
}
