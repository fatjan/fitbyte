package dto

import (
	"errors"
	"fmt"
	"time"

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
	DoneAtFrom        string `form:"doneAtFrom"`
	DoneAtTo          string `form:"doneAtTo"`
	CaloriesBurnedMin int    `form:"caloriesBurnedMin" binding:"gte=0"`
	CaloriesBurnedMax int    `form:"caloriesBurnedMax" binding:"gte=0"`
}

type ActivityPayload struct {
	Limit             int
	Offset            int
	ActivityType      string
	DoneAtFrom        time.Time
	DoneAtTo          time.Time
	CaloriesBurnedMin int
	CaloriesBurnedMax int
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

func (d *ActivityQueryParamRequest) ValidateActivityFilter() (*ActivityPayload, error) {
	payloadData := &ActivityPayload{}

	// Set default values for Limit and Offset
	if d.Limit <= 0 {
		payloadData.Limit = 5
	} else {
		payloadData.Limit = d.Limit
	}

	if d.Offset < 0 {
		payloadData.Offset = 0
	} else {
		payloadData.Offset = d.Offset
	}

	// Validate CaloriesBurnedMin and CaloriesBurnedMax
	if d.CaloriesBurnedMin > d.CaloriesBurnedMax {
		return nil, errors.New("caloriesBurnedMin cannot be greater than caloriesBurnedMax")
	}
	payloadData.CaloriesBurnedMin = d.CaloriesBurnedMin
	payloadData.CaloriesBurnedMax = d.CaloriesBurnedMax

	// Parse DoneAtFrom
	if d.DoneAtFrom != "" {
		parsedDate, err := time.Parse(time.RFC3339, d.DoneAtFrom)
		if err != nil {
			return nil, fmt.Errorf("invalid DoneAtFrom format: %w", err)
		}
		payloadData.DoneAtFrom = parsedDate
	}

	// Parse DoneAtTo
	if d.DoneAtTo != "" {
		parsedDate, err := time.Parse(time.RFC3339, d.DoneAtTo)
		if err != nil {
			return nil, fmt.Errorf("invalid DoneAtTo format: %w", err)
		}
		payloadData.DoneAtTo = parsedDate
	}

	// Validate DoneAtFrom and DoneAtTo
	if !payloadData.DoneAtFrom.IsZero() && !payloadData.DoneAtTo.IsZero() {
		if payloadData.DoneAtFrom.After(payloadData.DoneAtTo) {
			return nil, errors.New("doneAtFrom cannot be after doneAtTo")
		}
	}

	// Set ActivityType
	payloadData.ActivityType = d.ActivityType

	return payloadData, nil
}
