package validators

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

func New() *Validator {
	return &Validator{validate: validator.New()}
}

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) Struct(s any) error {
	return v.validate.Struct(s)
}

func (v *Validator) Validate(r *http.Request, data any) error {
	if err := v.Struct(data); err != nil {
		return err
	}
	return nil
}
