package handler

import (
	"payment/internal/broker"
	"payment/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services      *service.Service
	kafkaProducer *broker.KafkaProducer
}

func NewHandler(services *service.Service, kafkaProducer *broker.KafkaProducer) *Handler {
	return &Handler{
		services:      services,
		kafkaProducer: kafkaProducer}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/payment")
	{
		api.POST("/", h.createPayment)
		api.POST("/cancel", h.cancelPayment)
		api.GET("/:id", h.getById)
		api.GET("/", h.getAll)
		api.DELETE("/:id", h.deletePayment)
	}

	return router
}
