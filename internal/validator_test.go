package internal

import (
	"fmt"
	"testing"
)

func newValidatorSut(value string) *Validator {
	return &Validator{
		value:  value,
		Errors: &Errors{},
	}
}

func TestValidator_exactlyLengthValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		argument   string
		expected   string
	}{
		{"valid", "12345", "5", ""},
		{"boundary1", "12345", "4", "the length must be exactly 4"},
		{"boundary2", "12345", "6", "the length must be exactly 6"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.exactlyLength = tc.argument
		sut.exactlyLengthValidate()
		assert(t, tc.expected, sut.Errors, tc.value, tc.argument)
	}
}

func TestValidator_minLengthValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		argument   string
		expected   string
	}{
		{"valid1", "12345", "5", ""},
		{"valid2", "12345", "4", ""},
		{"boundary", "12345", "6", "the length must be no less than 6"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.minLength = tc.argument
		sut.minLengthValidate()
		assert(t, tc.expected, sut.Errors, tc.value, tc.argument)
	}
}

func TestValidator_maxLengthValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		argument   string
		expected   string
	}{
		{"valid1", "12345", "5", ""},
		{"valid2", "12345", "6", ""},
		{"boundary", "12345", "4", "the length must be no more than 4"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.maxLength = tc.argument
		sut.maxLengthValidate()
		assert(t, tc.expected, sut.Errors, tc.value, tc.argument)
	}
}

func TestValidator_notEmptyValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "abc", ""},
		{"valid2", "0", ""},
		{"valid3", "false", ""},
		{"valid4", "null", ""},
		{"valid5", " ", ""},
		{"invalid", "", "cannot be blank"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.notEmpty = true
		sut.notEmptyValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_digitValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid", "12345", ""},
		{"invalid", "abc12", "must contain digits only"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.digit = true
		sut.digitValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_alphaValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid", "abcABC", ""},
		{"invalid", "abcABC123", "must contain English letters only"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.alpha = true
		sut.alphaValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_alphanumericValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid", "abcABC123", ""},
		{"invalid", "abcABC123<>", "must contain English letters and digits only"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.alphanumeric = true
		sut.alphanumericValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_patternValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		argument   string
		expected   string
	}{
		{"valid1", "example-value", `^[\w+=,.@-]+$`, ""},
		{"valid2", "valid+=,.@-", `^[\w+=,.@-]+$`, ""},
		{"invalid", "invalid+=,.@-<>", `^[\w+=,.@-]+$`, "must be in a valid format"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.pattern = tc.argument
		sut.patternValidate()
		assert(t, tc.expected, sut.Errors, tc.value, tc.argument)
	}
}

func assert(t *testing.T, expected string, actual *Errors, value string, argument string) {
	if expected == "" {
		assertNoError(t, actual, value, argument)
	} else {
		assertError(t, expected, actual, value, argument)
	}
}

func assertNoError(t *testing.T, actual *Errors, value string, argument string) {
	if actual.HasError() {
		t.Errorf(formatMessage(NoError, actual, value, argument))
	}
}

func assertError(t *testing.T, expected string, actual *Errors, value string, argument string) {
	if !actual.HasError() {
		t.Errorf(formatMessage(expected, NoError, value, argument))
	} else if actual.Error() != expected {
		t.Errorf(formatMessage(expected, actual, value, argument))
	}
}

func formatMessage(expected string, actual any, value string, argument string) string {
	return fmt.Sprintf("expected: %s, actual: %+v, value: %s, argument: %s", expected, actual, value, argument)
}

const NoError = "<no error>"
const NoArgument = "<n/a>"
