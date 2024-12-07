package internal

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

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

	min            string
	max            string
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
	float          bool
	url            bool
	pattern        string
	enum           string
	timestamp      string
}

func (v *Validator) validate() error {
	v.minValidate()
	v.maxValidate()
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
	v.floatValidate()
	v.urlValidate()
	v.patternValidate()
	v.enumValidate()
	v.timestampValidate()

	if !v.HasError() {
		return nil
	}
	return fmt.Errorf(v.Errors.Error())
}

func (v *Validator) minValidate() {
	if v.min == "" {
		return
	}

	if value, err1 := strconv.ParseInt(v.value, 10, 64); err1 == nil {
		condition, err2 := strconv.ParseInt(v.min, 10, 64)
		if err2 != nil {
			v.AddArgumentError(fmt.Errorf("invalid min: %s", v.min))
			return
		}
		v.wrapAnyValidate(value, validation.Min(condition))
		return
	}

	if value, err1 := strconv.ParseFloat(v.value, 64); err1 == nil {
		condition, err2 := strconv.ParseFloat(v.min, 64)
		if err2 != nil {
			v.AddArgumentError(fmt.Errorf("invalid min: %s", v.min))
			return
		}
		v.wrapAnyValidate(value, validation.Min(condition))
		return
	}

	v.AddValidationError(fmt.Errorf("%s is not supported: %s", reflect.ValueOf(v.value).Type(), v.value))
}

func (v *Validator) maxValidate() {
	if v.max == "" {
		return
	}

	if value, err1 := strconv.ParseInt(v.value, 10, 64); err1 == nil {
		condition, err2 := strconv.ParseInt(v.max, 10, 64)
		if err2 != nil {
			v.AddArgumentError(fmt.Errorf("invalid max: %s", v.max))
			return
		}
		v.wrapAnyValidate(value, validation.Max(condition))
		return
	}

	if value, err1 := strconv.ParseFloat(v.value, 64); err1 == nil {
		condition, err2 := strconv.ParseFloat(v.max, 64)
		if err2 != nil {
			v.AddArgumentError(fmt.Errorf("invalid max: %s", v.max))
			return
		}
		v.wrapAnyValidate(value, validation.Max(condition))
		return
	}

	v.AddValidationError(fmt.Errorf("%s is not supported: %s", reflect.ValueOf(v.value).Type(), v.value))
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

func (v *Validator) floatValidate() {
	if !v.float {
		return
	}
	v.wrapValidate(is.Float)
}

func (v *Validator) urlValidate() {
	if !v.url {
		return
	}
	v.wrapValidate(is.RequestURL)
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

func (v *Validator) enumValidate() {
	if v.enum == "" {
		return
	}

	enumerations := strings.Split(v.enum, ",")
	if !slices.Contains(enumerations, v.value) {
		v.AddValidationError(fmt.Errorf("must be a valid value: %v", enumerations))
	}
}

func (v *Validator) timestampValidate() {
	if v.timestamp == "" {
		return
	}

	layouts := map[string]string{
		"rfc3339":  time.RFC3339,
		"datetime": time.DateTime,
		"date":     time.DateOnly,
		"time":     time.TimeOnly,
	}

	if layout, ok := layouts[strings.ToLower(v.timestamp)]; ok {
		v.wrapValidate(validation.Date(layout))
	} else {
		v.AddArgumentError(fmt.Errorf("not found layout: %s in [%v]", v.timestamp, layouts))
	}
}

func (v *Validator) wrapValidate(rules ...validation.Rule) {
	err := validation.Validate(v.value, rules...)
	if err != nil {
		v.AddValidationError(err)
	}
}

func (v *Validator) wrapAnyValidate(value any, rules ...validation.Rule) {
	err := validation.Validate(value, rules...)
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
