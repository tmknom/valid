package internal

import (
	"fmt"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func newValidator() *Validator {
	return &Validator{
		Errors: &Errors{},
	}
}

type Validator struct {
	value string
	*Errors

	exactlyLength string
	minLength     string
	maxLength     string
}

func (v *Validator) validate() error {
	v.exactlyLengthValidate()
	v.minLengthValidate()
	v.maxLengthValidate()

	if !v.HasError() {
		return nil
	}
	return fmt.Errorf(v.Errors.Error())
}

func (v *Validator) exactlyLengthValidate() {
	if v.exactlyLength == "" {
		return
	}
	if number, ok := v.toInt(v.exactlyLength); ok {
		v.wrapValidate(validation.Length(number, number))
	}
}

func (v *Validator) minLengthValidate() {
	if v.minLength == "" {
		return
	}
	if number, ok := v.toInt(v.minLength); ok {
		v.wrapValidate(validation.Length(number, 0))
	}
}

func (v *Validator) maxLengthValidate() {
	if v.maxLength == "" {
		return
	}
	if number, ok := v.toInt(v.maxLength); ok {
		v.wrapValidate(validation.Length(0, number))
	}
}

func (v *Validator) wrapValidate(rules ...validation.Rule) {
	err := validation.Validate(v.value, rules...)
	if err != nil {
		v.AddValidationError(err)
	}
}

func (v *Validator) toInt(s string) (int, bool) {
	val, err := strconv.Atoi(s)
	if err != nil {
		v.AddArgumentError(err)
		return 0, false
	}
	return val, true
}
