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
