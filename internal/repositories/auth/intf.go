package auth

import (
	"context"

	"github.com/fatjan/fitbyte/internal/models"
)

type Repository interface {
	FindByEmail(context.Context, string) (*models.User, error)
	Post(context.Context, *models.User) (int, error)
}
