package activity

import (
	"context"

	"github.com/fatjan/fitbyte/internal/models"
)

type Repository interface {
	Post(ctx context.Context, activity *models.Activity) (*models.Activity, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, activity *models.Activity) (*models.Activity, error)
}
