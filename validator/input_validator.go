package validator

import (
	"github.com/go-playground/validator/v10"
)

type InputValidator struct {
	Validator *validator.Validate
}

func (iv *InputValidator) Validate(i interface{}) error {
	if err := iv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
