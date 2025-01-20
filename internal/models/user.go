package models

import "time"

type PreferenceType string
type WeightUnitType string
type HeightUnitType string

const (
	Cardio	PreferenceType = "CARDIO"
	Weight 	PreferenceType = "WEIGHT"

	KG		WeightUnitType = "KG"
	LBS		WeightUnitType = "LBS"

	CM		HeightUnitType = "CM"
	INCH	HeightUnitType = "INCH"
)

type User struct {
	ID         int       		`db:"id"`
	Email      string    		`db:"email"`
	Password   string    		`db:"password_hash"`
	Name       string    		`db:"name"`
	Preference PreferenceType   `db:"preference"`
	WeightUnit WeightUnitType   `db:"weight_unit"`
	HeightUnit HeightUnitType   `db:"height_unit"`
	Weight     int    			`db:"weight"`
	Height     int    			`db:"height"`
	ImageUri   string    		`db:"image_uri"`
	CreatedAt  time.Time 		`db:"created_at"`
	UpdatedAt  time.Time 		`db:"updated_at"`
}
