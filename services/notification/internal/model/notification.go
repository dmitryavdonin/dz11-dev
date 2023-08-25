package model

import (
	"time"
)

type Notification struct {
	ID         int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	UserId     int       `gorm:"type:integer;not null" json:"user_id,omitempty"`
	OrderId    int       `gorm:"type:integer;not null" json:"order_id,omitempty"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt time.Time `gorm:"not null" json:"modified_at,omitempty"`
}

type NewNotification struct {
	UserId  int    `json:"user_id,omitempty"`
	OrderId int    `json:"order_id,omitempty"`
	Message string `json:"message"`
}
