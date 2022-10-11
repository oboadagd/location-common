package dto

import (
	"github.com/go-playground/validator/v10"
)

type GetByLatitudeLongitudeRangeRequest struct {
	LatitudeMin  float64 `json:"latitudeMin" validate:"required,min=-90,max=90,maxDecimals"`
	LatitudeMax  float64 `json:"latitudeMax" validate:"required,min=-90,max=90,maxDecimals"`
	LongitudeMin float64 `json:"longitudeMin" validate:"required,min=-180,max=180,maxDecimals"`
	LongitudeMax float64 `json:"longitudeMax" validate:"required,min=-180,max=180,maxDecimals"`
}

func (ur *GetByLatitudeLongitudeRangeRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
