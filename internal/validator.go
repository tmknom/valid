package internal

import (
	"fmt"
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
		Value: &Value{},
	}
}

type Value struct {
	raw  string
	mask bool
}

func (v *Value) Unmasked() string {
	return v.raw
}

func (v *Value) Masked() string {
	if v.mask {
		return MaskedValue
	}
	return v.raw
}

const MaskedValue = "***"

type Validator struct {
	*Value
	*Errors

	min            string
	max            string
	exactLength    string
	minLength      string
	maxLength      string
	notEmpty       bool
	digit          bool
	alpha          bool
	alphanumeric   bool
	ascii          bool
	printableASCII bool
	lowerCase      bool
	upperCase      bool
	int            bool
	float          bool
	url            bool
	domain         bool
	email          bool
	semver         bool
	uuid           bool
	base64         bool
	json           bool
	pattern        string
	enum           string
	timestamp      string
}

func (v *Validator) validate() error {
	v.Errors = &Errors{Value: v.Value}

	v.minValidate()
	v.maxValidate()
	v.exactLengthValidate()
	v.minLengthValidate()
	v.maxLengthValidate()
	v.notEmptyValidate()
	v.digitValidate()
	v.alphaValidate()
	v.alphanumericValidate()
	v.asciiValidate()
	v.printableASCIIValidate()
	v.lowerCaseValidate()
	v.upperCaseValidate()
	v.intValidate()
	v.floatValidate()
	v.urlValidate()
	v.domainValidate()
	v.emailValidate()
	v.semverValidate()
	v.uuidValidate()
	v.base64Validate()
	v.jsonValidate()
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

	if value, err1 := strconv.ParseInt(v.Value.Unmasked(), 10, 64); err1 == nil {
		condition, err2 := strconv.ParseInt(v.min, 10, 64)
		if err2 != nil {
			v.AddArgumentError(fmt.Errorf("--min must be an integer number"))
			return
		}
		v.wrapAnyValidate(value, validation.Min(condition))
		return
	}

	if value, err1 := strconv.ParseFloat(v.Value.Unmasked(), 64); err1 == nil {
		condition, err2 := strconv.ParseFloat(v.min, 64)
		if err2 != nil {
			v.AddArgumentError(fmt.Errorf("--min must be an float number"))
			return
		}
		v.wrapAnyValidate(value, validation.Min(condition))
		return
	}

	v.AddArgumentError(fmt.Errorf("--min cannot validate non-numeric value"))
}

func (v *Validator) maxValidate() {
	if v.max == "" {
		return
	}

	if value, err1 := strconv.ParseInt(v.Value.Unmasked(), 10, 64); err1 == nil {
		condition, err2 := strconv.ParseInt(v.max, 10, 64)
		if err2 != nil {
			v.AddArgumentError(fmt.Errorf("--max must be an integer number"))
			return
		}
		v.wrapAnyValidate(value, validation.Max(condition))
		return
	}

	if value, err1 := strconv.ParseFloat(v.Value.Unmasked(), 64); err1 == nil {
		condition, err2 := strconv.ParseFloat(v.max, 64)
		if err2 != nil {
			v.AddArgumentError(fmt.Errorf("--max must be an float number"))
			return
		}
		v.wrapAnyValidate(value, validation.Max(condition))
		return
	}
	v.AddArgumentError(fmt.Errorf("--max cannot validate non-numeric value"))
}

func (v *Validator) exactLengthValidate() {
	if v.exactLength == "" {
		return
	}

	number, err := strconv.Atoi(v.exactLength)
	if err != nil {
		v.AddArgumentError(fmt.Errorf("--exact-length must be an integer number"))
		return
	}
	v.wrapValidate(validation.Length(number, number))
}

func (v *Validator) minLengthValidate() {
	if v.minLength == "" {
		return
	}

	number, err := strconv.Atoi(v.minLength)
	if err != nil {
		v.AddArgumentError(fmt.Errorf("--min-length must be an integer number"))
		return
	}
	v.wrapValidate(validation.Length(number, 0))
}

func (v *Validator) maxLengthValidate() {
	if v.maxLength == "" {
		return
	}

	number, err := strconv.Atoi(v.maxLength)
	if err != nil {
		v.AddArgumentError(fmt.Errorf("--max-length must be an integer number"))
		return
	}
	v.wrapValidate(validation.Length(0, number))
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

func (v *Validator) lowerCaseValidate() {
	if !v.lowerCase {
		return
	}
	v.wrapValidate(is.LowerCase)
}

func (v *Validator) upperCaseValidate() {
	if !v.upperCase {
		return
	}
	v.wrapValidate(is.UpperCase)
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

func (v *Validator) domainValidate() {
	if !v.domain {
		return
	}
	v.wrapValidate(is.Domain)
}

func (v *Validator) emailValidate() {
	if !v.email {
		return
	}
	v.wrapValidate(is.EmailFormat)
}

func (v *Validator) semverValidate() {
	if !v.semver {
		return
	}
	v.wrapValidate(is.Semver)
}

func (v *Validator) uuidValidate() {
	if !v.uuid {
		return
	}
	v.wrapValidate(is.UUID)
}

func (v *Validator) base64Validate() {
	if !v.base64 {
		return
	}
	v.wrapValidate(is.Base64)
}

func (v *Validator) jsonValidate() {
	if !v.json {
		return
	}
	v.wrapValidate(is.JSON)
}

func (v *Validator) patternValidate() {
	if v.pattern == "" {
		return
	}

	regex, err := regexp.Compile(v.pattern)
	if err != nil {
		message := fmt.Sprintf("--pattern \"%s\" is not a valid regular expression", v.pattern)
		v.AddArgumentError(fmt.Errorf(message))
		return
	}
	v.wrapValidate(validation.Match(regex))
}

func (v *Validator) enumValidate() {
	if v.enum == "" {
		return
	}

	enumerations := strings.Split(v.enum, ",")
	if !slices.Contains(enumerations, v.Value.Unmasked()) {
		v.AddValidationError(fmt.Errorf("must be one of %v", enumerations))
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

	lowerTimestamp := strings.ToLower(v.timestamp)
	if layout, ok := layouts[lowerTimestamp]; ok {
		err := validation.Validate(v.Value.Unmasked(), validation.Date(layout))
		if err != nil {
			v.AddValidationError(fmt.Errorf("must be a valid %s", lowerTimestamp))
		}
	} else {
		var keys []string
		for key, _ := range layouts {
			keys = append(keys, key)
		}
		message := fmt.Sprintf("--timestamp must be one of %v", keys)
		v.AddArgumentError(fmt.Errorf(message))
	}
}

func (v *Validator) wrapValidate(rules ...validation.Rule) {
	err := validation.Validate(v.Value.Unmasked(), rules...)
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
