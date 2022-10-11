package dto

import (
	"github.com/go-playground/validator/v10"
)

type GetUsersByLocationAndRadiusRequest struct {
	Latitude   float64 `json:"latitude" validate:"required,min=-90,max=90,maxDecimals"`
	Longitude  float64 `json:"longitude" validate:"required,min=-180,max=180,maxDecimals"`
	Radius     float64 `json:"radius"`
	Page       uint64  `json:"page"`
	ItemsLimit uint64  `json:"itemsLimit"`
}

func (ur *GetUsersByLocationAndRadiusRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
