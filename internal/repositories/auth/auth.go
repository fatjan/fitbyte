package auth

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fatjan/fitbyte/internal/models"
	"github.com/fatjan/fitbyte/internal/pkg/exceptions"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type repository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Post(ctx context.Context, user *models.User) (int, error) {
	var newID int
	now := time.Now()

	query := `
			INSERT INTO users (email, password_hash, created_at, updated_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		user.Email,
		user.Password,
		now,
		now,
	).Scan(&newID)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == pq.ErrorCode("23505") {
				return 0, exceptions.ErrConflict
			}
		}
		return 0, err
	}

	return newID, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}

	err := r.db.QueryRowContext(ctx,
		"SELECT id, email, password_hash FROM users WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	return user, nil
}
