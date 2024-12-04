package internal

import (
	"fmt"
	"regexp"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func newValidator() *Validator {
	return &Validator{
		Errors: &Errors{},
	}
}

type Validator struct {
	value string
	*Errors

	exactlyLength  string
	minLength      string
	maxLength      string
	notEmpty       bool
	digit          bool
	alpha          bool
	alphanumeric   bool
	ascii          bool
	printableASCII bool
	int            bool
	pattern        string
}

func (v *Validator) validate() error {
	v.exactlyLengthValidate()
	v.minLengthValidate()
	v.maxLengthValidate()
	v.notEmptyValidate()
	v.digitValidate()
	v.alphaValidate()
	v.alphanumericValidate()
	v.asciiValidate()
	v.printableASCIIValidate()
	v.intValidate()
	v.patternValidate()

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

func (v *Validator) notEmptyValidate() {
	if !v.notEmpty {
		return
	}
	v.wrapValidate(validation.Required)
}

func (v *Validator) digitValidate() {
	if !v.digit {
		return
	}
	v.wrapValidate(is.Digit)
}

func (v *Validator) alphaValidate() {
	if !v.alpha {
		return
	}
	v.wrapValidate(is.Alpha)
}

func (v *Validator) alphanumericValidate() {
	if !v.alphanumeric {
		return
	}
	v.wrapValidate(is.Alphanumeric)
}

func (v *Validator) asciiValidate() {
	if !v.ascii {
		return
	}
	v.wrapValidate(is.ASCII)
}

func (v *Validator) printableASCIIValidate() {
	if !v.printableASCII {
		return
	}
	v.wrapValidate(is.PrintableASCII)
}

func (v *Validator) intValidate() {
	if !v.int {
		return
	}
	v.wrapValidate(is.Int)
}

func (v *Validator) patternValidate() {
	if v.pattern == "" {
		return
	}

	regex, err := regexp.Compile(v.pattern)
	if err != nil {
		v.AddArgumentError(err)
		return
	}
	v.wrapValidate(validation.Match(regex))
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
