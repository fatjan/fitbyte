package user

import (
	"context"
	"github.com/fatjan/fitbyte/internal/dto"
	"github.com/fatjan/fitbyte/internal/models"
)

type Repository interface {
	GetUser(ctx context.Context, id int) (*models.User, error)
	Update(context.Context, int, *dto.UserPatchRequest) error
}