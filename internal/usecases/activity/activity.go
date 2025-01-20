package activity

import (
	"context"
	"strconv"

	"github.com/fatjan/fitbyte/internal/dto"
	"github.com/fatjan/fitbyte/internal/models"
	"github.com/fatjan/fitbyte/internal/pkg/exceptions"
	"github.com/fatjan/fitbyte/internal/repositories/activity"
)

type useCase struct {
	activityRepository activity.Repository
}

func NewUseCase(activityRepository activity.Repository) UseCase {
	return &useCase{
		activityRepository,
	}
}

func (u *useCase) PostActivity(ctx context.Context, activity *dto.ActivityRequest, userId int) (*dto.ActivityResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	newActivity := &models.Activity{
		UserId:         userId,
		ActivityType:   activity.ActivityType,
		DoneAt:         activity.DoneAt,
		DurationInMin:  activity.DurationInMinutes,
		CaloriesBurned: activity.ActivityType.GetTotalCalories(activity.DurationInMinutes),
	}

	res, err := u.activityRepository.Post(ctx, newActivity)
	if err != nil {
		return nil, err
	}

	activityId := strconv.Itoa(res.ID)

	return &dto.ActivityResponse{
		ActivityId:        activityId,
		ActivityType:      res.ActivityType,
		DoneAt:            res.DoneAt,
		DurationInMinutes: res.DurationInMin,
		CaloriesBurned:    res.CaloriesBurned,
		CreatedAt:         res.CreatedAt,
		UpdatedAt:         res.UpdatedAt,
	}, nil
}

func (u *useCase) DeleteActivity(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if id == "" {
		return exceptions.ErrNotFound
	}

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt <= 0 {
		return exceptions.ErrNotFound
	}

	err = u.activityRepository.Delete(ctx, idInt)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return err
		}

		return err
	}

	return nil
}

func (u *useCase) UpdateActivity(ctx context.Context, activity *dto.ActivityRequest, userID int, activityID string) (*dto.ActivityResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if activityID == "" {
		return nil, exceptions.ErrNotFound
	}

	activityIdInt, err := strconv.Atoi(activityID)
	if err != nil {
		return nil, err
	}

	newActivity := &models.Activity{
		ActivityType:   activity.ActivityType,
		DoneAt:         activity.DoneAt,
		DurationInMin:  activity.DurationInMinutes,
		CaloriesBurned: activity.ActivityType.GetTotalCalories(activity.DurationInMinutes),
		ID:             activityIdInt,
		UserId:         userID,
	}

	res, err := u.activityRepository.Update(ctx, newActivity)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return nil, exceptions.ErrNotFound
		}
		return nil, err
	}

	activityId := strconv.Itoa(res.ID)

	return &dto.ActivityResponse{
		ActivityId:        activityId,
		ActivityType:      res.ActivityType,
		DoneAt:            res.DoneAt,
		DurationInMinutes: res.DurationInMin,
		CaloriesBurned:    res.CaloriesBurned,
		CreatedAt:         res.CreatedAt,
		UpdatedAt:         res.UpdatedAt,
	}, nil
}
