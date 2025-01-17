package duck

import (
	"context"

	"github.com/fatjan/fitbyte/internal/models"
)

type Repository interface {
	GetDucks(ctx context.Context) ([]*models.Duck, error)
	GetDuckByID(ctx context.Context, id int) (*models.Duck, error)
}
