package dto

import (
	"github.com/go-playground/validator/v10"
)

// GetLastByUserNameRequest is a request of GetLastByUserName method
type GetLastByUserNameRequest struct {
	UserName string `json:"username" validate:"required,min=4,max=16,patternazAZ09"` // username. It is required, belongs to length range 4 to 16, belongs regex pattern PatternUserNameRegexString
}

// Validate applies validations specified previously
func (ur *GetLastByUserNameRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
