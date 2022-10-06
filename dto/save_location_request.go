package dto

import (
	"github.com/go-playground/validator/v10"
	"math"
	"regexp"
)

type SaveLocationRequest struct {
	UserName  string  `json:"username" validate:"required,min=4,max=16,patterazAZ09"`
	Latitude  float64 `json:"latitude" validate:"required,min=-90,max=90,maxDecimals"`
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180,maxDecimals"`
}

func IsMaxDecimals(fl validator.FieldLevel) bool {
	valuef := fl.Field().Float() * math.Pow(10.0, float64(8))
	println(valuef)
	extra := valuef - float64(int(valuef))

	return extra == 0
}

func IsPatternUserName(fl validator.FieldLevel) bool {
	PatternUserNameRegexString := "^[a-zA-Z0-9]+$"
	PatternUserNameRegex := regexp.MustCompile(PatternUserNameRegexString)

	return PatternUserNameRegex.MatchString(fl.Field().String())
}

type CustomValidatorSaveLoc struct {
	Validator *validator.Validate
}

func (cv *CustomValidatorSaveLoc) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}
