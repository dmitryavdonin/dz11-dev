package repository

import (
	"notification/internal/model"

	"gorm.io/gorm"
)

type NotificationPostgres struct {
	db *gorm.DB
}

func NewNotificationPostgres(db *gorm.DB) *NotificationPostgres {
	return &NotificationPostgres{db: db}
}

func (r *NotificationPostgres) Create(n model.Notification) (int, error) {
	result := r.db.Create(&n)
	if result.Error != nil {
		return 0, result.Error
	}
	return n.ID, nil
}

func (r *NotificationPostgres) GetById(orderId int) (model.Notification, error) {
	var item model.Notification
	result := r.db.First(&item, "order_id = ?", orderId)
	return item, result.Error
}

func (r *NotificationPostgres) GetAll(limit int, offset int) ([]model.Notification, error) {
	var items []model.Notification
	result := r.db.Limit(limit).Offset(offset).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, result.Error
}

func (r *NotificationPostgres) Delete(orderId int) error {
	result := r.db.Delete(&model.Notification{}, "order_id = ?", orderId)
	return result.Error
}
