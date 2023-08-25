package service

import (
	"notification/internal/model"
	"notification/internal/repository"
)

type NotificationService struct {
	repo repository.Notification
}

func NewNotificationService(repo repository.Notification) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) Create(input model.Notification) (int, error) {
	return s.repo.Create(input)
}

func (s *NotificationService) GetById(order_id int) ([]model.Notification, error) {
	return s.repo.GetById(order_id)
}

func (s *NotificationService) GetAll(limit int, offset int) ([]model.Notification, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *NotificationService) Delete(order_id int) error {
	return s.repo.Delete(order_id)
}
