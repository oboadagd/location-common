package dto

import (
	"github.com/go-playground/validator/v10"
	"math"
	"regexp"
)

// SaveLocationRequest is a http request of Save service.
type SaveLocationRequest struct {
	UserName  string  `json:"username" validate:"required,min=4,max=16,patternazAZ09"`    // username located in one geographic point. It is required, belongs to length range 4 to 16, belongs regex pattern PatternUserNameRegexString
	Latitude  float64 `json:"latitude" validate:"required,min=-90,max=90,maxDecimals"`    // latitude coordinate of username's location. It is required, belongs to range -90 to 90, allows 8 decimal positions
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180,maxDecimals"` // longitude coordinate of username's location. It is required, belongs to range -180 to 180, allows 8 decimal positions
}

// IsMaxDecimals is a validation rule. It defines the maximum decimal positions
// of a float number
func IsMaxDecimals(fl validator.FieldLevel) bool {
	valuef := fl.Field().Float() * math.Pow(10.0, float64(8))
	extra := valuef - float64(int(valuef))

	return extra == 0
}

// IsPatternUserName is a validation rule. It defines that input string
// matches to regex pattern PatternUserNameRegexString
func IsPatternUserName(fl validator.FieldLevel) bool {
	PatternUserNameRegexString := "^[a-zA-Z0-9]+$"
	PatternUserNameRegex := regexp.MustCompile(PatternUserNameRegexString)

	return PatternUserNameRegex.MatchString(fl.Field().String())
}

// CustomValidatorSaveLoc is custom validator.
type CustomValidatorSaveLoc struct {
	Validator *validator.Validate // validator
}

// Validate applies validations specified previously
func (cv *CustomValidatorSaveLoc) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}
