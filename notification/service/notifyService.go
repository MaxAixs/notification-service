package service

import (
	"fmt"
	"github.com/google/uuid"
	"notification-service/notification"
	"notification-service/notification/repository"
	"notification-service/pkg/mailgun"
	"sync"
	"time"
)

type NotificationService struct {
	repo    *repository.Repository
	mailgun mailgun.Mailer
}

func NewNotificationService(repo *repository.Repository, mg mailgun.Mailer) *NotificationService {
	return &NotificationService{repo: repo, mailgun: mg}
}

func (n *NotificationService) ProcessDeadlineData(deadlineData []notification.DeadlineUserInfo) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(deadlineData))

	for _, data := range deadlineData {
		wg.Add(1)

		go func(data notification.DeadlineUserInfo) {
			defer wg.Done()

			if err := n.processDeadlineItem(data); err != nil {
				errChan <- err
			}
		}(data)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return err
	}

	return nil
}

func (n *NotificationService) processDeadlineItem(data notification.DeadlineUserInfo) error {
	msg := createNotify(data)

	if err := n.mailgun.SendEmail(msg.Email, msg.Topic, msg.Body); err != nil {
		return fmt.Errorf("failed to send email to %s: %w", msg.Email, err)
	}

	msg.Status = "sent"
	now := time.Now()
	msg.SentAt = &now

	return n.repo.CreateNotification(msg)
}

func createNotify(data notification.DeadlineUserInfo) notification.Notification {
	return notification.Notification{
		ID:        uuid.New(),
		UserId:    data.UserId,
		Email:     data.Email,
		ItemId:    data.ItemId,
		Topic:     "Reminder: Task due date is approaching",
		Body:      fmt.Sprintf("Hey! Time is running out to complete your task №%d: %s ⏳ Hurry up!", data.ItemId, data.Description),
		Status:    "not_sent",
		CreatedAt: time.Now(),
	}
}
