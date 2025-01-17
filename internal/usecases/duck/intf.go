package duck

import (
	"context"

	"github.com/fatjan/fitbyte/internal/dto"
)

type Usecase interface {
	GetDucks(ctx context.Context) ([]*dto.Duck, error)
	GetDuckByID(ctx context.Context, id string) (*dto.Duck, error)
}
