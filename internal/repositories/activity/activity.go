package activity

import (
	"context"

	"github.com/fatjan/fitbyte/internal/models"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Post(ctx context.Context, activity *models.Activity) (*models.Activity, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	query := `
		INSERT INTO activities (user_id, activity_type, done_at, duration_in_minutes, calories_burned)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		activity.UserId,
		activity.ActivityType,
		activity.DoneAt,
		activity.DurationInMin,
		activity.CaloriesBurned,
	).Scan(&activity.ID, &activity.CreatedAt, &activity.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return activity, nil
}
