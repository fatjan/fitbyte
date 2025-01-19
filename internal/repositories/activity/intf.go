package activity

import (
	"context"

	"github.com/fatjan/fitbyte/internal/models"
)

type Repository interface {
	Post(ctx context.Context, activity *models.Activity) (*models.Activity, error)
}
