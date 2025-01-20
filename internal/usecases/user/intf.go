package user

import (
	"context"
	"github.com/fatjan/fitbyte/internal/dto"
)

type UseCase interface {
	GetUser(*dto.UserRequest) (*dto.User, error)
	UpdateUser(context.Context, int, *dto.UserPatchRequest) (*dto.UserPatchResponse, error)
}