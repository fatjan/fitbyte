package auth

import (
	"context"

	"github.com/fatjan/fitbyte/internal/config"
	"github.com/fatjan/fitbyte/internal/dto"
	"github.com/fatjan/fitbyte/internal/models"
	"github.com/fatjan/fitbyte/internal/pkg/exceptions"
	"github.com/fatjan/fitbyte/internal/pkg/jwt_helper"
	"github.com/fatjan/fitbyte/internal/repositories/auth"
)

type useCase struct {
	authRepository auth.Repository
	cfg            *config.Config
}

func NewUseCase(authRepository auth.Repository, cfg *config.Config) UseCase {
	return &useCase{
		authRepository: authRepository,
		cfg:            cfg,
	}
}

func (uc *useCase) Login(ctx context.Context, authRequest *dto.AuthRequest) (*dto.AuthResponse, error) {
	err := authRequest.ValidatePayloadAuth()
	if err != nil {
		return nil, err
	}

	user, err := uc.authRepository.FindByEmail(ctx, authRequest.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, exceptions.ErrNotFound
	}

	err = authRequest.ComparePassword(user.Password)
	if err != nil {
		return nil, err
	}

	token, err := jwt_helper.SignJwt(uc.cfg.JwtKey, user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Email: authRequest.Email,
		Token: token,
	}, nil
}

func (uc *useCase) Register(ctx context.Context, authRequest *dto.AuthRequest) (*dto.AuthResponse, error) {
	err := authRequest.ValidatePayloadAuth()
	if err != nil {
		return nil, err
	}

	err = authRequest.HashPassword()
	if err != nil {
		return nil, err
	}

	newAuth := &models.User{
		Email:    authRequest.Email,
		Password: authRequest.Password,
	}

	id, err := uc.authRepository.Post(ctx, newAuth)
	if err != nil {
		return nil, err
	}

	token, err := jwt_helper.SignJwt(uc.cfg.JwtKey, id)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Email: authRequest.Email,
		Token: token,
	}, nil
}
