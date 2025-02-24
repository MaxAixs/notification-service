package service

import (
	"notification-service/notification"
	"notification-service/notification/repository"
	"notification-service/pkg/analyticService"
	"notification-service/pkg/mailgun"
)

type DeadlineNotifier interface {
	ProcessDeadlineData(deadlineData []notification.DeadlineUserInfo) error
	ProcessDeadlineItem(data notification.DeadlineUserInfo) error
}

type CompletedTaskNotifier interface {
	ProcessDoneTasksData(tasks []analyticService.CompletedTasks) error
	ProcessDoneTask(task analyticService.CompletedTasks) error
}

type NotificationService struct {
	DeadlineNotifier
	CompletedTaskNotifier
}

func NewNotificationService(db *repository.Repository, mailgun mailgun.Mailer) *NotificationService {
	return &NotificationService{
		DeadlineNotifier:      NewDeadlineNotifierService(db, mailgun),
		CompletedTaskNotifier: NewDoneTaskNotifier(db, mailgun),
	}
}
