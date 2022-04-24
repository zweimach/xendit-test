package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func NewError() ErrorResponse {
	e := ErrorResponse{}
	e.Errors = make([]string, 0, 5)
	return e
}

func (e *ErrorResponse) AddError(err error) {
	e.Errors = append(e.Errors, err.Error())
}

func (e *ErrorResponse) AddNotFound(name string) {
	e.Errors = append(e.Errors, fmt.Sprintf("%v not found", name))
}

func (e *ErrorResponse) AddBadRequest() {
	e.Errors = append(e.Errors, "invalid request")
}

func (e *ErrorResponse) AddValidationError(err error) {
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.Errors = append(e.Errors, fmt.Sprintf("validation for '%v' failed: %v", strings.ToLower(v.Field()), v.Tag()))
	}
}
