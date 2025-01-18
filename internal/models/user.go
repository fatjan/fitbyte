package models

import (
	"time"

	"github.com/fatjan/fitbyte/internal/dto"
)

type User struct {
	ID         int       `db:"id"`
	Email      string    `db:"email"`
	Password   string    `db:"password_hash"`
	Name       string    `db:"name"`
	Preference dto.PreferenceType    `db:"preference"`
	WeightUnit dto.WeightUnitType    `db:"weight_unit"`
	HeightUnit dto.HeightUnitType    `db:"height_unit"`
	Weight     int    `db:"weight"`
	Height     int    `db:"height"`
	ImageUri   string    `db:"image_uri"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
