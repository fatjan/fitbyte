package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/fatjan/fitbyte/internal/models"
)

var validate = validator.New()

func (u *UserPatchRequest) Validate() error {
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

type User struct {
	Preference      models.PreferenceType 	`json:"preference" validate:"oneof=CARDIO WEIGHT"`
	WeightUnit      models.WeightUnitType 	`json:"weightUnit" validate:"oneof=KG LBS"`
	HeightUnit      models.HeightUnitType	`json:"heightUnit" validate:"oneof=CM INCH"`
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
	Preference *models.PreferenceType `json:"preference" validate:"required,oneof=CARDIO WEIGHT"`
	WeightUnit *models.WeightUnitType `json:"weightUnit" validate:"required,oneof=KG LBS"`
	HeightUnit *models.HeightUnitType `json:"heightUnit" validate:"required,oneof=CM INCH"`
	Weight     *int            `json:"weight" validate:"required,min=10,max=1000"`
	Height     *int            `json:"height" validate:"required,min=3,max=250"`
	Name       *string         `json:"name" validate:"min=2,max=60"`
	ImageUri   *string         `json:"imageUri" validate:"url"`
}

type UserPatchResponse struct {
	Preference      *models.PreferenceType 	`json:"preference" validate:"oneof=CARDIO WEIGHT"`
	WeightUnit      *models.WeightUnitType 	`json:"weightUnit" validate:"oneof=KG LBS"`
	HeightUnit      *models.HeightUnitType	`json:"heightUnit" validate:"oneof=CM INCH"`
	Weight      	*int 			`json:"weight"`
	Height      	*int 			`json:"height"`
	Name            string 			`json:"name"`
	ImageUri    	string 			`json:"imageUri"`
}