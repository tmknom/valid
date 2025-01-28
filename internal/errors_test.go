package internal

import (
	"fmt"
	"testing"
)

func TestErrors_HasError(t *testing.T) {
	cases := []struct {
		annotation string
		validation bool
		argument   bool
		expected   bool
	}{
		{"has-validation-error", true, false, true},
		{"has-argument-error", false, true, true},
		{"has-validation-and-argument-error", true, true, true},
		{"no-error", false, false, false},
	}

	for _, tc := range cases {
		sut := &Errors{value: &Value{raw: "pseudo"}}
		if tc.validation {
			sut.AddValidationError(fmt.Errorf("validation error"))
		}
		if tc.argument {
			sut.AddArgumentError(fmt.Errorf("argument error"))
		}
		actual := sut.HasError()

		format := "\n annotation: %s\n expected:   %v\n actual:     %v\n validation: %v\n argument:   %v"
		if tc.expected != actual {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, actual, tc.validation, tc.argument))
		}
	}
}

func TestErrors_Error(t *testing.T) {
	cases := []struct {
		annotation  string
		validations []string
		arguments   []string
		expected    string
	}{
		{"one-validation-error", []string{"one"}, nil, "Validation error: The specified value \"pseudo\" is invalid. Issues: one."},
		{"multi-validation-error", []string{"one", "two"}, nil, "Validation error: The specified value \"pseudo\" is invalid. Issues: one, two."},
		{"one-argument-error", nil, []string{"one"}, "Argument error: one."},
		{"multi-argument-error", nil, []string{"one", "two"}, "Argument error: one, two."},
		{"complex-error", []string{"one", "two"}, []string{"three", "four"}, "Validation error: The specified value \"pseudo\" is invalid. Issues: one, two; Argument error: three, four."},
		{"no-error", nil, nil, ""},
	}

	for _, tc := range cases {
		sut := &Errors{value: &Value{raw: "pseudo"}}
		for _, validation := range tc.validations {
			sut.AddValidationError(fmt.Errorf(validation))
		}
		for _, argument := range tc.arguments {
			sut.AddArgumentError(fmt.Errorf(argument))
		}
		actual := sut.Error()

		format := "\n annotation:  %s\n expected:    %v\n actual:      %v\n validations: %v\n arguments:   %v"
		if tc.expected != actual {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, actual, tc.validations, tc.arguments))
		}
	}
}
