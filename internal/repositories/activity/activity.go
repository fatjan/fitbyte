package activity

import (
	"context"
	"database/sql"

	"github.com/fatjan/fitbyte/internal/models"
	"github.com/fatjan/fitbyte/internal/pkg/exceptions"
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

func (r *repository) Delete(ctx context.Context, id int) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	query := `DELETE FROM activities WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return exceptions.ErrNotFound
	}

	return nil
}

func (r *repository) Update(ctx context.Context, activity *models.Activity) (*models.Activity, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	query := `
		UPDATE activities
		SET activity_type = $1, done_at = $2, duration_in_minutes = $3, calories_burned = $4, updated_at = now()
		WHERE id = $5 and user_id = $6
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		activity.ActivityType,
		activity.DoneAt,
		activity.DurationInMin,
		activity.CaloriesBurned,
		activity.ID,
		activity.UserId,
	).Scan(&activity.ID, &activity.CreatedAt, &activity.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exceptions.ErrNotFound
		}
		return nil, err
	}

	return activity, nil
}
