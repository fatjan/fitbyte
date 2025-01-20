package activity

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fatjan/fitbyte/internal/dto"
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

func (r *repository) Get(ctx context.Context, activity *dto.ActivityQueryParamRequest) ([]*dto.ActivityResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	var (
		query  = `SELECT id, activity_type, done_at, duration_in_minutes, calories_burned, created_at FROM activities WHERE 1=1`
		params []interface{}
		idx    = 1
	)

	// Add filters dynamically
	if activity.ActivityType != "" {
		query += ` AND activity_type = $` + fmt.Sprint(idx)
		params = append(params, activity.ActivityType)
		idx++
	}

	if activity.DoneAtFrom != "" {
		query += ` AND done_at >= $` + fmt.Sprint(idx)
		params = append(params, activity.DoneAtFrom)
		idx++
	}

	if activity.DoneAtTo != "" {
		query += ` AND done_at <= $` + fmt.Sprint(idx)
		params = append(params, activity.DoneAtTo)
		idx++
	}

	if activity.CaloriesBurnedMin > 0 {
		query += ` AND calories_burned >= $` + fmt.Sprint(idx)
		params = append(params, activity.CaloriesBurnedMin)
		idx++
	}

	if activity.CaloriesBurnedMax > 0 {
		query += ` AND calories_burned <= $` + fmt.Sprint(idx)
		params = append(params, activity.CaloriesBurnedMax)
		idx++
	}

	// Add pagination
	query += ` LIMIT $` + fmt.Sprint(idx)
	params = append(params, activity.Limit)
	idx++

	query += ` OFFSET $` + fmt.Sprint(idx)
	params = append(params, activity.Offset)

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse the result
	var results []*dto.ActivityResponse
	for rows.Next() {
		var (
			id              string
			activityType    models.ActivityType
			doneAt          string
			durationMinutes int
			caloriesBurned  int
			createdAt       string
			updatedAt       string
		)

		err := rows.Scan(&id, &activityType, &doneAt, &durationMinutes, &caloriesBurned, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		// Append result to slice
		results = append(results, &dto.ActivityResponse{
			ActivityId:        id,
			ActivityType:      activityType,
			DoneAt:            doneAt,
			DurationInMinutes: durationMinutes,
			CaloriesBurned:    caloriesBurned,
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
		})
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
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
