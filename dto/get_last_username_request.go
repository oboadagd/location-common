package dto

import (
	"github.com/go-playground/validator/v10"
)

type GetLastByUserNameRequest struct {
	UserName string `json:"username" validate:"required,min=4,max=16,patterazAZ09""`
}

func (ur *GetLastByUserNameRequest) Validate() error {
	if err := validator.New().Struct(ur); err != nil {
		return err
	}

	return nil
}
