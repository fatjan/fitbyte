package activity

import (
	"context"

	"github.com/fatjan/fitbyte/internal/dto"
	"github.com/fatjan/fitbyte/internal/models"
)

type Repository interface {
	Post(ctx context.Context, activity *models.Activity) (*models.Activity, error)
	Get(ctx context.Context, activity *dto.ActivityPayload, userId int) ([]*dto.ActivityResponse, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, activity *models.Activity) (*models.Activity, error)
}
