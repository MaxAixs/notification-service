package client

import "notification-service/pkg/analyticService"

type AnalyticsClient interface {
	GetDoneTasksAnalytic() ([]analyticService.CompletedTasks, error)
	CloseConnection()
}
