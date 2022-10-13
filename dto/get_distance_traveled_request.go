package dto

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// GetDistanceTraveledRequest is a http request of GetDistanceTraveledRequest service.
type GetDistanceTraveledRequest struct {
	UserName    string    `json:"username" validate:"required,min=4,max=16,patternazAZ09"` // username located in one geographic point. It is required, belongs to length range 4 to 16, belongs regex pattern PatternUserNameRegexString
	InitialDate time.Time `json:"initialDate"`                                             // initial date range to accumulate traveled distance
	FinalDate   time.Time `json:"finalDate"`                                               // final date range to accumulate traveled distance
}

// Validate applies validations specified previously
func (ur *GetDistanceTraveledRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
