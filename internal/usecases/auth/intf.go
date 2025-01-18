package auth

import (
	"context"

	"github.com/fatjan/fitbyte/internal/dto"
)

type UseCase interface {
	Login(context.Context, *dto.AuthRequest) (*dto.AuthResponse, error)
	Register(context.Context, *dto.AuthRequest) (*dto.AuthResponse, error)
}
