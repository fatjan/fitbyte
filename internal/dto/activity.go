package dto

import (
	"errors"

	"github.com/fatjan/fitbyte/internal/models"
)

type ActivityRequest struct {
	ActivityType      models.ActivityType `validate:"required,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAt            string              `validate:"required,iso8601"`
	DurationInMinutes int                 `validate:"required,gte=1"`
}

type ActivityQueryParamRequest struct {
	Limit             int    `form:"limit" binding:"gte=0"`
	Offset            int    `form:"offset" binding:"gte=0"`
	ActivityType      string `form:"activityType" binding:"omitempty,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAtFrom        string `form:"doneAtFrom" binding:"omitempty,datetime=2006-01-02"`
	DoneAtTo          string `form:"doneAtTo" binding:"omitempty,datetime=2006-01-02"`
	CaloriesBurnedMin int    `form:"caloriesBurnedMin" binding:"gte=0"`
	CaloriesBurnedMax int    `form:"caloriesBurnedMax" binding:"gte=0"`
}

type ActivityResponse struct {
	ActivityId        string              `json:"activityId"`
	ActivityType      models.ActivityType `json:"activityType"`
	DoneAt            string              `json:"doneAt"`
	DurationInMinutes int                 `json:"durationInMinutes"`
	CaloriesBurned    int                 `json:"caloriesBurned"`
	CreatedAt         string              `json:"createdAt"`
	UpdatedAt         string              `json:"updatedAt"`
}

func (d *ActivityQueryParamRequest) ValidateActivityFilter() error {
	if d.Limit == 0 {
		d.Limit = 5
	}
	if d.Offset == 0 {
		d.Offset = 0
	}

	if d.CaloriesBurnedMin > d.CaloriesBurnedMax {
		return errors.New("caloriesBurnedMin cannot be greater than caloriesBurnedMax")
	}

	return nil
}
