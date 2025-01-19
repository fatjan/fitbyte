package dto

import "github.com/fatjan/fitbyte/internal/models"

type ActivityRequest struct {
	ActivityType      models.ActivityType `validate:"required,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAt            string              `validate:"required,iso8601"`
	DurationInMinutes int                 `validate:"required,gte=1"`
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
