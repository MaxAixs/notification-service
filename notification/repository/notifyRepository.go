package repository

import (
	"database/sql"
	"fmt"
	"notification-service/notification"
)

type NotifyRepository struct {
	db *sql.DB
}

func NewNotifyRepository(db *sql.DB) *NotifyRepository {
	return &NotifyRepository{db: db}
}

func (n *NotifyRepository) CreateNotification(msg notification.Notification) error {
	q := `INSERT INTO notifications (id,user_id,email,item_id,topic,body,status,created_at,sent_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := n.db.Exec(q, msg.ID, msg.UserId, msg.Email, msg.ItemId, msg.Topic, msg.Body, msg.Status, msg.CreatedAt, msg.SentAt)
	if err != nil {
		return fmt.Errorf("cannot insert notification: %w", err)
	}

	return nil
}
