package dto

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (u *UserPatchRequest) Validate() error {
	// Register the custom URI validation
	validate.RegisterValidation("uri", uriCustomValidate)

	err := validate.Struct(u)
	if err != nil {
		// Handle validation errors here
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "email":
				return fmt.Errorf("field '%s' must be a valid email address", err.Field())
			case "min":
				return fmt.Errorf("field '%s' must be at least %s characters long", err.Field(), err.Param())
			case "max":
				return fmt.Errorf("field '%s' cannot exceed %s characters", err.Field(), err.Param())
			case "uri":
				return fmt.Errorf("field '%s' must be a valid URI", err.Field())
			default:
				return fmt.Errorf("field '%s' failed validation on '%s' tag", err.Field(), err.Tag())
			}
		}
	}
	return nil
}

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
	Preference      PreferenceType 	`json:"preference" validate:"oneof=CARDIO WEIGHT"`
	WeightUnit      WeightUnitType 	`json:"weightUnit" validate:"oneof=KG LBS"`
	HeightUnit      HeightUnitType	`json:"heightUnit" validate:"oneof=CM INCH"`
	Weight      	int 			`json:"weight"`
	Height      	int 			`json:"height"`
	Email           string 			`json:"email"`
	Name            string 			`json:"name"`
	ImageUri    	string 			`json:"imageUri"`
}

type UserRequest struct {
	UserID int `json:"id"`
}

type UserPatchRequest struct {
	Preference      *string `json:"preference" validate:"required"`
	WeightUnit      *string `json:"weight_unit" validate:"required"`
	HeightUnit    	*string `json:"height_unit" validate:"required"`
	Weight      	*string `json:"weight" validate:"required"`
	Height      	*string `json:"height" validate:"required"`
	Name     		*string `json:"name" validate:"min=2,max=60"`
	ImageUri 		*string `json:"imageUri" validate:"url"`
}

// Create a custom URI validation function
func uriCustomValidate(fl validator.FieldLevel) bool {
	uri := fl.Field().String()
	if uri == "" {
		return true // It's valid if the field is empty (omitempty).
	}

	// Try to parse the URI
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return false // URI is not valid
	}

	// Add custom validation logic https or http:
	if !strings.HasPrefix(uri, "http://") && !strings.HasPrefix(uri, "https://") {
		return false // Custom check to make sure the URI starts with http:// or https://
	}

	return true
}