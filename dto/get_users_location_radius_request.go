package dto

import (
	"github.com/go-playground/validator/v10"
)

// GetUsersByLocationAndRadiusRequest is a http request of GetUsersByLocationAndRadius service
type GetUsersByLocationAndRadiusRequest struct {
	Latitude   float64 `json:"latitude" validate:"min=-90,max=90,maxDecimals"`    // latitude coordinate of username's location. It belongs to range -90 to 90, allows 8 decimal positions
	Longitude  float64 `json:"longitude" validate:"min=-180,max=180,maxDecimals"` // longitude coordinate of username's location. It belongs to range -90 to 90, allows 8 decimal positions
	Radius     float64 `json:"radius" validate:"required,gt=0"`                   // range radius. It is required, belongs to range (0 to +infinite)
	Page       uint64  `json:"page" validate:"min=1"`                             // page number to show up. It belongs to range [1 to +infinite)
	ItemsLimit uint64  `json:"itemsLimit" validate:"min=1"`                       // quantity of items per page. It belongs to range [1 to +infinite)
}

// Validate applies validations specified previously
func (ur *GetUsersByLocationAndRadiusRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
