package user

import (
	"context"
	"fmt"
	"github.com/fatjan/fitbyte/internal/dto"
	"github.com/fatjan/fitbyte/internal/repositories/user"
	"log"
)

type useCase struct {
	userRepository user.Repository
}

func NewUseCase(userRepository user.Repository) UseCase {
	return &useCase{userRepository: userRepository}
}

func (u *useCase) GetUser(userRequest *dto.UserRequest) (*dto.User, error) {
	user, err := u.userRepository.GetUser(userRequest.UserID)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		Preference:		 user.Preference,
		WeightUnit:		 user.WeightUnit,
		HeightUnit:		 user.HeightUnit,
		Weight:		 	 user.Weight,
		Height:		 	 user.Height,
		Email:           user.Email,
		Name:            user.Name,
		ImageUri:    	 user.ImageUri,
	}, nil
}
func (u *useCase) UpdateUser(ctx context.Context, userID int, request *dto.UserPatchRequest) (*dto.UserPatchResponse, error) {
	// Get existing user
	user, err := u.userRepository.GetUser(userID)
	if err != nil {
		log.Println(fmt.Errorf("failed to get user: %w", err))
		return nil, err
	}

	// Create update request with current values as defaults
	updateRequest := &dto.UserPatchRequest{}

	// Update fields if provided in request
	if request != nil {
		// check key request name
		if request.Name != nil {
			updateRequest.Name = request.Name
			user.Name = *request.Name
		}

		// check key request ImageUri
		if request.ImageUri != nil {
			updateRequest.ImageUri = request.ImageUri
			user.ImageUri = *request.ImageUri
		}
	}

	updateRequest.Preference = request.Preference
	updateRequest.WeightUnit = request.WeightUnit
	updateRequest.HeightUnit = request.HeightUnit
	updateRequest.Weight = request.Weight
	updateRequest.Height = request.Height

	// Update user in repository
	if err = u.userRepository.Update(ctx, userID, updateRequest); err != nil {
		log.Println(fmt.Errorf("failed to update user: %w", err))
		return nil, err
	}

	// Return new user data
	return &dto.UserPatchResponse{
		Preference:      updateRequest.Preference,
		WeightUnit:		 updateRequest.WeightUnit,
		HeightUnit:		 updateRequest.HeightUnit,
		Weight:		 	 updateRequest.Weight,
		Height:		 	 updateRequest.Height,
		Name:            user.Name,
		ImageUri:    	 user.ImageUri,
	}, nil
}