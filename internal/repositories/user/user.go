package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/fatjan/fitbyte/internal/dto"
	"log"
	"strings"

	"github.com/fatjan/fitbyte/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	PG_DUPLICATE_ERROR = "23505"
)

type repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUser(id int) (*models.User, error) {
	nullFields := struct {
		Preference      sql.NullString
		WeightUnit      sql.NullString
		HeightUnit      sql.NullString
		Weight      	sql.NullInt64
		Height      	sql.NullInt64
		Name            sql.NullString
		ImageUri    	sql.NullString
	}{}

	user := &models.User{}

	err := r.db.QueryRow(`
	select 
		email, name, preference, weight_unit, height_unit, weight, height, image_uri 
	from users 
	where id = $1;`, id).Scan(
		&user.Email,
		&nullFields.Name,
		&nullFields.Preference,
		&nullFields.WeightUnit,
		&nullFields.HeightUnit,
		&nullFields.Weight,
		&nullFields.Height,
		&nullFields.ImageUri,
	)

	if err != nil {
		return nil, err
	}

	if nullFields.Preference.Valid {
		preference := nullFields.Preference.String
		if preference == string(models.Cardio) || preference == string(models.Weight) {
			user.Preference = models.PreferenceType(preference)
		} else {
			user.Preference = "" 
		}
	} else {
		user.Preference = "" 
	}

	if nullFields.WeightUnit.Valid {
		weightUnit := nullFields.WeightUnit.String
		if weightUnit == string(models.KG) || weightUnit == string(models.LBS) {
			user.WeightUnit = models.WeightUnitType(weightUnit)
		} else {
			user.WeightUnit = "" 
		}
	} else {
		user.WeightUnit = "" 
	}

	if nullFields.HeightUnit.Valid {
		heightUnit := nullFields.HeightUnit.String
		if heightUnit == string(models.CM) || heightUnit == string(models.INCH) {
			user.HeightUnit = models.HeightUnitType(heightUnit)
		} else {
			user.HeightUnit = "" 
		}
	} else {
		user.HeightUnit = "" 
	}

	if nullFields.Weight.Valid {
		user.Weight = int(nullFields.Weight.Int64)
	} else {
		user.Weight = 0
	}
	
	if nullFields.Height.Valid {
		user.Height = int(nullFields.Height.Int64)
	} else {
		user.Height = 0
	}

	user.Name = nullFields.Name.String
	user.ImageUri = nullFields.ImageUri.String

	return user, nil
}

func (r *repository) Update(ctx context.Context, userID int, request *dto.UserPatchRequest) error {
	baseQuery := `UPDATE users SET `
	var setClauses []string
	var args []interface{}
	var argIndex int = 1

	if request != nil {
		if request.Preference != nil {
			setClauses = append(setClauses, fmt.Sprintf(`preference = $%d`, argIndex))
			args = append(args, *request.Preference)
			argIndex++
		}

		if request.WeightUnit != nil {
			setClauses = append(setClauses, fmt.Sprintf(`weight_unit = $%d`, argIndex))
			args = append(args, *request.WeightUnit)
			argIndex++
		}

		if request.HeightUnit != nil {
			setClauses = append(setClauses, fmt.Sprintf(`height_unit = $%d`, argIndex))
			args = append(args, *request.HeightUnit)
			argIndex++
		}

		if request.Weight != nil {
			setClauses = append(setClauses, fmt.Sprintf(`weight = $%d`, argIndex))
			args = append(args, *request.Weight)
			argIndex++
		}

		if request.Height != nil {
			setClauses = append(setClauses, fmt.Sprintf(`height = $%d`, argIndex))
			args = append(args, *request.Height)
			argIndex++
		}

		if request.Name != nil {
			setClauses = append(setClauses, fmt.Sprintf(`name = $%d`, argIndex))
			args = append(args, *request.Name)
			argIndex++
		}

		if request.ImageUri != nil {
			setClauses = append(setClauses, fmt.Sprintf(`image_uri = $%d`, argIndex))
			args = append(args, *request.ImageUri)
			argIndex++
		}
	}

	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}

	baseQuery += strings.Join(setClauses, ", ")
	baseQuery += fmt.Sprintf(` WHERE id = $%d`, argIndex)
	args = append(args, userID)

	result, err := r.db.ExecContext(ctx, baseQuery, args...)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == PG_DUPLICATE_ERROR {
			return fmt.Errorf("duplicate email")
		}
		
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("error query")
		return err
	}
	if rowsAffected == 0 {
		log.Println("failed update user")
		return errors.New("update query failed")
	}

	return nil
}