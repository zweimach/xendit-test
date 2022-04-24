package utils

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator(v *validator.Validate) *CustomValidator {
	return &CustomValidator{v}
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		err := err.(validator.ValidationErrors)
		return err
	}
	return nil
}
