package dto

import (
	"github.com/go-playground/validator/v10"
)

// CreateLocationHistoryRequest is a request of Create method.
type CreateLocationHistoryRequest struct {
	UserName  string  `json:"username" validate:"required,min=4,max=16,patternazAZ09"`    // username located in one geographic point. It is required, belongs to length range 4 to 16, belongs regex pattern PatternUserNameRegexString
	Latitude  float64 `json:"latitude" validate:"required,min=-90,max=90,maxDecimals"`    // latitude coordinate of username's location. It is required, belongs to range -90 to 90, allows 8 decimal positions
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180,maxDecimals"` // longitude coordinate of username's location. It is required, belongs to range -180 to 180, allows 8 decimal positions
	Distance  float64 `json:"distance"`                                                   //traveled distance by username from last to current location
}

// Validate applies validations specified previously
func (cl *CreateLocationHistoryRequest) Validate() error {
	if err := validator.New().Struct(cl); err != nil {
		return err
	}

	return nil
}
