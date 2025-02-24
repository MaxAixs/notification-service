package analyticService

import "github.com/google/uuid"

type CompletedTasks struct {
	UserId uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	Count  int32     `json:"count"`
}
