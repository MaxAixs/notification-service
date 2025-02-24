package worker

import (
	"context"
	"fmt"
	"log"
	"notification-service/notification/service"
	"notification-service/pkg/analyticService/client"
	"time"
)

type AnalyticWorker struct {
	service    *service.NotificationService
	gRPCClient client.AnalyticsClient
}

func NewAnalyticWorker(service *service.NotificationService, analyticClient client.AnalyticsClient) *AnalyticWorker {
	return &AnalyticWorker{
		service:    service,
		gRPCClient: analyticClient,
	}
}

func (a *AnalyticWorker) Start(ctx context.Context) error {
	ticker := time.NewTicker(7 * 24 * time.Hour)
	defer ticker.Stop()

	log.Println("AnalyticWorker started")

	for {
		select {
		case <-ctx.Done():
			log.Println("AnalyticWorker shutting down")
			return ctx.Err()

		case <-ticker.C:
			res, err := a.gRPCClient.GetDoneTasksAnalytic()
			if err != nil {
				return fmt.Errorf("cant get analytics request: %v", err)
			}

			log.Printf("Received analytics data: %+v", res)

			if err := a.service.ProcessDoneTasksData(res); err != nil {
				return fmt.Errorf("process done tasks data error: %v", err)
			}
		}
	}
}
