package repository

import (
	"database/sql"
	"notification-service/notification"
)

type Notification interface {
	CreateNotification(msg notification.Notification) error
}

type Repository struct {
	Notification
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Notification: NewNotifyRepository(db),
	}
}
