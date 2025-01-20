package activity

import (
	"context"

	"github.com/fatjan/fitbyte/internal/dto"
)

type UseCase interface {
	PostActivity(context.Context, *dto.ActivityRequest, int) (*dto.ActivityResponse, error)
	DeleteActivity(ctx context.Context, id string) error
	UpdateActivity(ctx context.Context, activity *dto.ActivityRequest, userID int, activityID string) (*dto.ActivityResponse, error)
}
