package internal

import (
	"fmt"
	"testing"
)

func newValidatorSut(value string) *Validator {
	return &Validator{
		value:  value,
		Errors: &Errors{value: value},
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
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.max = tc.argument
		sut.maxValidate()
		assert(t, tc.expected, sut.Errors, tc.value, tc.argument)
	}
}

func TestValidator_exactLengthValidate(t *testing.T) {
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
		sut.exactLength = tc.argument
		sut.exactLengthValidate()
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

func TestValidator_lowerCaseValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "abc", ""},
		{"valid2", "abc123<>", ""},
		{"invalid", "abcABC", "must be in lower case"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.lowerCase = true
		sut.lowerCaseValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_upperCaseValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid1", "ABC", ""},
		{"valid2", "ABC123<>", ""},
		{"invalid", "abcABC", "must be in upper case"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.upperCase = true
		sut.upperCaseValidate()
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

func TestValidator_domainValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid", "example.com", ""},
		{"invalid", "localhost", "must be a valid domain"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.domain = true
		sut.domainValidate()
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

func TestValidator_uuidValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid", "b4563933-7bf8-4d9b-b4cd-d0c6f85b5925", ""},
		{"invalid", "b4563933-7bf8-4d9b-b4cd-d0c6f85b592", "must be a valid UUID"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.uuid = true
		sut.uuidValidate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_base64Validate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid", "dmFsaWQ=", ""},
		{"invalid", "invalid", "must be encoded in Base64"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.base64 = true
		sut.base64Validate()
		assert(t, tc.expected, sut.Errors, tc.value, NoArgument)
	}
}

func TestValidator_jsonValidate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"valid", "[1, 2]", ""},
		{"invalid", "[1, 2,]", "must be in valid JSON format"},
	}

	for _, tc := range cases {
		sut := newValidatorSut(tc.value)
		sut.json = true
		sut.jsonValidate()
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
		{"invalid", "invalid", "foo,bar,baz", "must specify [foo bar baz]"},
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
		{"invalid1", "2024-08-09 12:34:56", "RFC3339", "must be a valid rfc3339"},
		{"invalid2", "2024-08-09T12:34:56Z", "datetime", "must be a valid datetime"},
		{"invalid3", "12:34:56", "date", "must be a valid date"},
		{"invalid4", "2024-08-09", "time", "must be a valid time"},
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
	expectedMessage := fmt.Sprintf("Validation error: The value \"%s\" is invalid. Issues: %s", value, expected)
	if !actual.HasError() {
		t.Errorf(formatMessage(expected, NoError, value, argument))
	} else if actual.Error() != expectedMessage {
		t.Errorf(formatMessage(expectedMessage, actual, value, argument))
	}
}

func formatMessage(expected string, actual any, value string, argument string) string {
	return fmt.Sprintf("\n expected: %s\n actual:   %+v\n value:    %s\n argument: %s", expected, actual, value, argument)
}

const NoError = "<no error>"
const NoArgument = "<n/a>"
