package handler

import (
	"net/http"
	"strconv"
	"time"

	"notification/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// create order
func (h *Handler) create(c *gin.Context) {
	logrus.Print("create(): BEGIN ")
	var input model.NewNotification
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("create(): Cannot parse input, error = %s", err.Error())
		return
	}

	now := time.Now()

	item := model.Notification{
		UserId:     input.UserId,
		OrderId:    input.OrderId,
		Message:    input.Message,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	logrus.Printf("create(): Try to create notification record: user_id = %d, order_id = %d, message = %s",
		item.UserId, item.OrderId, item.Message)
	id, err := h.services.Notification.Create(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("create(): Cannot create record, error = %s", err.Error())
		return
	}

	item.ID = id

	c.JSON(http.StatusOK, item)

	logrus.Print("create(): END")
}

// get order by id
func (h *Handler) getByOrderId(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.services.Notification.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// get all orders
func (h *Handler) getAll(c *gin.Context) {

	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var items []model.Notification
	items, err := h.services.Notification.GetAll(intLimit, offset)
	if err != nil {
		c.JSON(http.StatusBadGateway,
			StatusResponse{
				Status: "failed",
				Reason: err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(items), "data": items})
}

// delete order by id
func (h *Handler) deleteByOrderId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.Notification.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
