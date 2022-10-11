package dto

import (
	"github.com/go-playground/validator/v10"
)

type CreateLocationHistoryRequest struct {
	UserName  string  `json:"username" validate:"required,min=4,max=16,patternazAZ09""`
	Latitude  float64 `json:"latitude" validate:"required,min=-90,max=90,maxDecimals"`
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180,maxDecimals"`
	Distance  float64 `json:"distance"`
}

func (cl *CreateLocationHistoryRequest) Validate() error {
	if err := validator.New().Struct(cl); err != nil {
		return err
	}

	return nil
}
