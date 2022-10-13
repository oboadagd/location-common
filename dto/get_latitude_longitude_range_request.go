package dto

import (
	"github.com/go-playground/validator/v10"
)

// GetByLatitudeLongitudeRangeRequest is a request of GetByLatitudeLongitudeRange method
type GetByLatitudeLongitudeRangeRequest struct {
	LatitudeMin  float64 `json:"latitudeMin" validate:"required,min=-90,max=90,maxDecimals"`    // maximum latitude coordinate of username's location. It is required, belongs to range -90 to 90, allows 8 decimal positions
	LatitudeMax  float64 `json:"latitudeMax" validate:"required,min=-90,max=90,maxDecimals"`    // minimum latitude coordinate of username's location. It is required, belongs to range -90 to 90, allows 8 decimal positions
	LongitudeMin float64 `json:"longitudeMin" validate:"required,min=-180,max=180,maxDecimals"` // maximum longitude coordinate of username's location. It is required, belongs to range -90 to 90, allows 8 decimal positions
	LongitudeMax float64 `json:"longitudeMax" validate:"required,min=-180,max=180,maxDecimals"` // minimum longitude coordinate of username's location. It is required, belongs to range -90 to 90, allows 8 decimal positions
}

// Validate applies validations specified previously
func (ur *GetByLatitudeLongitudeRangeRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
