package duck

import (
	"context"

	"github.com/fatjan/fitbyte/internal/models"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewDuckRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetDucks(ctx context.Context) ([]*models.Duck, error) {
	query := "SELECT id, name FROM ducks"
	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ducks []*models.Duck
	for rows.Next() {
		var duck models.Duck
		if err := rows.StructScan(&duck); err != nil {
			return nil, err
		}
		ducks = append(ducks, &duck)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ducks, nil
}

func (r *repository) GetDuckByID(ctx context.Context, id int) (*models.Duck, error) {
	query := "SELECT id, name FROM ducks WHERE id = $1"
	var duck models.Duck
	if err := r.db.GetContext(ctx, &duck, query, id); err != nil {
		return nil, err
	}

	return &duck, nil
}
