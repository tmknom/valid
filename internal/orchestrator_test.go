package internal

import (
	"fmt"
	"testing"
)

func TestOrchestrator_orchestrate(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		mask       bool
		expected   string
	}{
		{"valid-masked", "valid-masked", true, ""},
		{"valid-unmasked", "valid-unmasked", false, ""},
		{"invalid-masked", "InvalidMasked", true, "Error: Validation error: The value \"***\" is invalid. Issues: must be in lower case"},
		{"invalid-unmasked", "InvalidUnmasked", false, "Error: Validation error: The value \"InvalidUnmasked\" is invalid. Issues: must be in lower case"},
	}

	for _, tc := range cases {
		sut := &Orchestrator{
			Value:     &Value{raw: tc.value, mask: tc.mask},
			Validator: &Validator{Errors: &Errors{}, lowerCase: true},
			Formatter: &Formatter{},
		}
		err := sut.orchestrate()

		format := "\n annotation: %s\n expected:   %s\n actual:     %+v\n value:      %s\n mask:       %v"
		if tc.expected == "" && err != nil {
			t.Errorf(fmt.Sprintf(format, tc.annotation, NoError, err.Error(), tc.value, tc.mask))
		}

		if tc.expected != "" {
			if err == nil {
				t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, NoError, tc.value, tc.mask))
			} else if err.Error() != tc.expected {
				t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, err.Error(), tc.value, tc.mask))
			}
		}
	}
}

func TestValue_Unmasked(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		mask       bool
		expected   string
	}{
		{"masked", "test-value", true, "test-value"},
		{"unmasked", "test-value", false, "test-value"},
	}

	for _, tc := range cases {
		sut := &Value{raw: tc.value, mask: tc.mask}
		actual := sut.Unmasked()

		format := "\n annotation: %s\n expected:   %s\n actual:     %+v\n value:      %s\n mask:       %v"
		if tc.expected != actual {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, actual, tc.value, tc.mask))
		}
	}
}

func TestValue_Masked(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		mask       bool
		expected   string
	}{
		{"masked", "test-value", true, "***"},
		{"unmasked", "test-value", false, "test-value"},
	}

	for _, tc := range cases {
		sut := &Value{raw: tc.value, mask: tc.mask}
		actual := sut.Masked()

		format := "\n annotation: %s\n expected:   %s\n actual:     %+v\n value:      %s\n mask:       %v"
		if tc.expected != actual {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, actual, tc.value, tc.mask))
		}
	}
}
