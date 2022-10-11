package dto

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type GetDistanceTraveledRequest struct {
	UserName    string    `json:"username" validate:"required,min=4,max=16,patterazAZ09""`
	InitialDate time.Time `json:"initialDate"`
	FinalDate   time.Time `json:"finalDate"`
}

func (ur *GetDistanceTraveledRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
