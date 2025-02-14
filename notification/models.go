package notification

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	ID        uuid.UUID  `json:"id"`
	UserId    uuid.UUID  `json:"user_id"`
	Email     string     `json:"email"`
	ItemId    int        `json:"item_id"`
	Topic     string     `json:"topic"`
	Body      string     `json:"message"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	SentAt    *time.Time `json:"sent_at"`
}

type DeadlineUserInfo struct {
	UserId      uuid.UUID `json:"user_id"`
	Email       string    `json:"email"`
	ItemId      int       `json:"item_id"`
	Description string    `json:"description"`
}
