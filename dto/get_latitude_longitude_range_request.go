package dto

import (
	"github.com/go-playground/validator/v10"
)

type GetByLatitudeLongitudeRangeRequest struct {
	LatitudeMin  float64 `json:"latitudeMin"`
	LatitudeMax  float64 `json:"latitudeMax"`
	LongitudeMin float64 `json:"longitudeMin"`
	LongitudeMax float64 `json:"longitudeMax"`
}

func (ur *GetByLatitudeLongitudeRangeRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
