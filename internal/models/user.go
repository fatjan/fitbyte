package models

import "time"

type User struct {
	ID              int       `db:"id"`
	Email           string    `db:"email"`
	Password        string    `db:"password_hash"`
	Name            string    `db:"name"`
	Preference     	string    `db:"preference"`
	WeightUnit 		string    `db:"weight_unit"`
	HeightUnit 		string    `db:"height_unit"`
	Weight 			string    `db:"weight"`
	Height 			string    `db:"height"`
	ImageUri    	string    `db:"image_uri"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}