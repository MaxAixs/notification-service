package client

import (
	"fmt"
	"github.com/MaxAixs/protos/gen/api/gen/api"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"notification-service/pkg/analyticService"
	"time"
)

type AnalyticClient struct {
	client api.AnalyticsDataClient
	conn   *grpc.ClientConn
}

func NewAnalyticsClient(port string) (*AnalyticClient, error) {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("NewNotifyClient error: %v", err)
	}

	grpcClient := api.NewAnalyticsDataClient(conn)

	return &AnalyticClient{client: grpcClient, conn: conn}, nil
}

func (a *AnalyticClient) GetDoneTasksAnalytic() ([]analyticService.CompletedTasks, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := a.client.FetchWeeklyCompletedTask(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weekly completed tasks from Analytics Service: %w", err)
	}

	doneTasksModel, err := convertToModel(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to convert weekly completed tasks from Analytics Service: %w", err)
	}

	return doneTasksModel, nil

}

func convertToModel(resp *api.WeeklyCompletedTasksResponse) ([]analyticService.CompletedTasks, error) {
	var doneTaskModel []analyticService.CompletedTasks

	if resp == nil {
		return nil, fmt.Errorf("resp is nil")
	}

	for _, task := range resp.Tasks {
		parsedUserId, err := uuid.Parse(task.UserId)
		if err != nil {
			return nil, fmt.Errorf("failed to parse task user id from model: %v", err)
		}

		doneTaskModel = append(doneTaskModel, analyticService.CompletedTasks{
			UserId: parsedUserId,
			Email:  task.Email,
			Count:  task.Count,
		})
	}

	return doneTaskModel, nil
}

func (a *AnalyticClient) CloseConnection() {
	a.conn.Close()
}
