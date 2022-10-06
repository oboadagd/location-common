package dto

import (
	"github.com/go-playground/validator/v10"
)

type GetUsersByLocationAndRadiusRequest struct {
	Latitude  float64 `json:"latitude" validate:`
	Longitude float64 `json:"longitude" validate:`
	Radius    float64 `json:"radius" validate:"required"`
}

func (ur *GetUsersByLocationAndRadiusRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
