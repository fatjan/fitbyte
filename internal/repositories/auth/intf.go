package auth

import (
	"github.com/fatjan/fitbyte/internal/models"
)

type Repository interface {
	FindByEmail(email string) (*models.User, error)
	Post(payload *models.User) (int, error)
}