package service

import (
	"notification/internal/model"
	"notification/internal/repository"
)

type Notification interface {
	Create(order model.Notification) (int, error)
	GetById(orderId int) (model.Notification, error)
	GetAll(limit int, offset int) ([]model.Notification, error)
	Delete(orderId int) error
}

type Services struct {
	Notification
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Notification: NewNotificationService(repos.Notification),
	}
}
