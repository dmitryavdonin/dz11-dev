package repository

import (
	"notification/internal/model"

	"gorm.io/gorm"
)

type Notification interface {
	Create(n model.Notification) (int, error)
	GetById(orderId int) (model.Notification, error)
	GetAll(limit int, offset int) ([]model.Notification, error)
	Delete(orderId int) error
}

type Repository struct {
	Notification
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Notification: NewNotificationPostgres(db),
	}
}
