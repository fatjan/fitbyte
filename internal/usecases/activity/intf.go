package activity

import (
	"context"

	"github.com/fatjan/fitbyte/internal/dto"
)

type UseCase interface {
	PostActivity(context.Context, *dto.ActivityRequest, int) (*dto.ActivityResponse, error)
}
