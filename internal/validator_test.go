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

func TestValidator_minValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		argument   string
		expected   string
	}{
		{"valid1", "10", "9", ""},
		{"valid2", "9", "9", ""},
		{"valid3", "-1", "-1", ""},
		{"valid4", "1.2", "1.1", ""},
		{"valid5", "1.1", "1.1", ""},
		{"valid6", "-1.1", "-1.1", ""},
		{"invalid1", "8", "9", "must be no less than 9"},
		{"invalid2", "-2", "-1", "must be no less than -1"},
		{"invalid3", "8.1", "9.1", "must be no less than 9.1"},
		{"invalid4", "-2.1", "-1.1", "must be no less than -1.1"},
		{"invalid5", "1", "a", "invalid min: a"},
		{"invalid6", "1.1", "a", "invalid min: a"},
		{"invalid7", "a", "9", "string is not supported: a"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.min = tc.argument
		sut.minValidate()
		assert(t, tc.expected, sut.Errors, tc.value, tc.argument)
	}
}

func TestValidator_maxValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		argument   string
		expected   string
	}{
		{"valid1", "9", "10", ""},
		{"valid2", "9", "9", ""},
		{"valid3", "-1", "-1", ""},
		{"valid4", "1.1", "1.2", ""},
		{"valid5", "1.1", "1.1", ""},
		{"valid6", "-1.1", "-1.1", ""},
		{"invalid1", "9", "8", "must be no greater than 8"},
		{"invalid2", "-1", "-2", "must be no greater than -2"},
		{"invalid3", "9.1", "8.1", "must be no greater than 8.1"},
		{"invalid4", "-1.1", "-2.1", "must be no greater than -2.1"},
		{"invalid5", "1", "a", "invalid max: a"},
		{"invalid6", "1.1", "a", "invalid max: a"},
		{"invalid7", "a", "9", "string is not supported: a"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.max = tc.argument
		sut.maxValidate()
		assert(t, tc.expected, sut.Errors, tc.value, tc.argument)
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

func TestValidator_asciiValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "abcABC123<>", ""},
		{"valid2", "'newline\r\ntab\t'", ""},
		{"valid3", "'\x00 ASCII \x7F'", ""},
		{"valid4", "'\x20 printable ASCII \x7E'", ""},
		{"valid5", "'\x19 not printable ASCII \x7F'", ""},
		{"invalid", "abcABC123<>あ", "must contain ASCII characters only"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.ascii = true
		sut.asciiValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_printableASCIIValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "abcABC123<>", ""},
		{"valid2", "'\x20 printable ASCII \x7E'", ""},
		{"invalid1", "abcABC123<>あ", "must contain printable ASCII characters only"},
		{"invalid2", "'newline\r\ntab\t'", "must contain printable ASCII characters only"},
		{"invalid3", "'\x00 ASCII \x7F'", "must contain printable ASCII characters only"},
		{"invalid4", "'\x19 not printable ASCII \x7F'", "must contain printable ASCII characters only"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.printableASCII = true
		sut.printableASCIIValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_intValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "12345", ""},
		{"valid2", "+12345", ""},
		{"valid3", "-12345", ""},
		{"invalid1", "abc123", "must be an integer number"},
		{"invalid2", "1.2", "must be an integer number"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.int = true
		sut.intValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_floatValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "12.345", ""},
		{"valid2", "12345", ""},
		{"invalid", "abc123", "must be a floating point number"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.float = true
		sut.floatValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_urlValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "https://example.com", ""},
		{"valid2", "https://localhost:8080/test.html", ""},
		{"invalid", "example.com", "must be a valid request URL"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.url = true
		sut.urlValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_emailValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "foo@example.com", ""},
		{"valid2", "foo+bar@example.com", ""},
		{"invalid", "example.com", "must be a valid email address"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.email = true
		sut.emailValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_semverValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "1.2.3", ""},
		{"valid2", "v1.2.3", ""},
		{"valid3", "v1.2.3-beta", ""},
		{"invalid", "1.2", "must be a valid semantic version"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.semver = true
		sut.semverValidate()
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

func TestValidator_enumValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		argument   string
		expected   string
	}{
		{"valid1", "bar", "foo,bar,baz", ""},
		{"valid2", "foo", "foo", ""},
		{"invalid", "invalid", "foo,bar,baz", "must be a valid value: [foo bar baz]"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.enum = tc.argument
		sut.enumValidate()
		assert(t, tc.expected, sut.Errors, tc.value, tc.argument)
	}
}

func TestValidator_timestampValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		argument   string
		expected   string
	}{
		{"valid1", "2024-08-09T12:34:56+07:00", "rfc3339", ""},
		{"valid2", "2024-08-09T12:34:56Z", "rfc3339", ""},
		{"valid3", "2024-08-09T12:34:56Z", "RFC3339", ""},
		{"valid4", "2024-08-09 12:34:56", "datetime", ""},
		{"valid5", "2024-08-09", "date", ""},
		{"valid6", "12:34:56", "time", ""},
		{"invalid1", "2024-08-09 12:34:56", "rfc3339", "must be a valid date"},
		{"invalid2", "2024-08-09T12:34:56Z", "datetime", "must be a valid date"},
		{"invalid3", "12:34:56", "date", "must be a valid date"},
		{"invalid4", "2024-08-09", "time", "must be a valid date"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.timestamp = tc.argument
		sut.timestampValidate()
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
