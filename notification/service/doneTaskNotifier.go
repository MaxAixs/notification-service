package service

import (
	"fmt"
	"github.com/google/uuid"
	"notification-service/notification"
	"notification-service/notification/repository"
	"notification-service/pkg/analyticService"
	"notification-service/pkg/mailgun"
	"sync"
	"time"
)

type DoneTaskNotifier struct {
	repo    *repository.Repository
	mailgun mailgun.Mailer
}

func NewDoneTaskNotifier(repo *repository.Repository, mg mailgun.Mailer) *DoneTaskNotifier {
	return &DoneTaskNotifier{repo: repo, mailgun: mg}
}

func (d *DoneTaskNotifier) ProcessDoneTasksData(tasks []analyticService.CompletedTasks) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(tasks))

	for _, task := range tasks {
		wg.Add(1)

		go func(task analyticService.CompletedTasks) {
			defer wg.Done()

			if err := d.ProcessDoneTask(task); err != nil {
				errChan <- err
			}
		}(task)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return err
	}

	return nil
}

func (d *DoneTaskNotifier) ProcessDoneTask(task analyticService.CompletedTasks) error {
	msg := createDoneTaskNotify(task)

	if err := d.mailgun.SendEmail(msg.Email, msg.Topic, msg.Body); err != nil {
		return fmt.Errorf("failed to send email to %s: %w", msg.Email, err)
	}

	msg.Status = "sent"
	now := time.Now()
	msg.SentAt = &now

	return d.repo.CreateNotification(msg)

}

func createDoneTaskNotify(task analyticService.CompletedTasks) notification.Notification {
	return notification.Notification{
		ID:        uuid.New(),
		UserId:    task.UserId,
		Email:     task.Email,
		Topic:     "Your Weekly Task Completion Summary 📊",
		Body:      fmt.Sprintf("Great job this week! You have completed %d tasks", task.Count),
		Status:    "not_sent",
		CreatedAt: time.Now(),
	}
}
