package internal

import (
	"fmt"
	"testing"
)

func TestOrchestrator_Orchestrate_Valid(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
	}{
		{"valid", "valid"},
	}

	for _, tc := range cases {
		sut := &Orchestrator{
			Value:     &Value{raw: tc.value},
			Validator: &Validator{Errors: &Errors{}, lowerCase: true},
			Formatter: &Formatter{},
		}
		err := sut.Orchestrate()

		format := "\n annotation: %s\n expected:   %s\n actual:     %+v\n value:      %s"
		if err != nil {
			t.Errorf(fmt.Sprintf(format, tc.annotation, NoError, err.Error(), tc.value))
		}
	}
}

func TestOrchestrator_Orchestrate_NamedValue(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		name       string
		expected   string
	}{
		{"specified-name", "Invalid", "test-id", "Error: Validation error: The specified test-id \"Invalid\" is invalid. Issues: must be in lower case"},
		{"not-specified-name", "Invalid", "", "Error: Validation error: The specified value \"Invalid\" is invalid. Issues: must be in lower case"},
	}

	for _, tc := range cases {
		sut := &Orchestrator{
			Value:     &Value{raw: tc.value, name: tc.name},
			Validator: &Validator{Errors: &Errors{}, lowerCase: true},
			Formatter: &Formatter{},
		}
		err := sut.Orchestrate()

		format := "\n annotation: %s\n expected:   %s\n actual:     %+v\n value:      %s\n name:       %v"
		if err == nil {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, NoError, tc.value, tc.name))
		} else if err.Error() != tc.expected {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, err.Error(), tc.value, tc.name))
		}
	}
}

func TestOrchestrator_Orchestrate_MaskedValue(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		mask       bool
		expected   string
	}{
		{"masked", "Invalid", true, "Error: Validation error: The specified value \"***\" is invalid. Issues: must be in lower case"},
		{"unmasked", "Invalid", false, "Error: Validation error: The specified value \"Invalid\" is invalid. Issues: must be in lower case"},
	}

	for _, tc := range cases {
		sut := &Orchestrator{
			Value:     &Value{raw: tc.value, mask: tc.mask},
			Validator: &Validator{Errors: &Errors{}, lowerCase: true},
			Formatter: &Formatter{},
		}
		err := sut.Orchestrate()

		format := "\n annotation: %s\n expected:   %s\n actual:     %+v\n value:      %s\n mask:       %v"
		if err == nil {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, NoError, tc.value, tc.mask))
		} else if err.Error() != tc.expected {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, err.Error(), tc.value, tc.mask))
		}
	}
}

func TestOrchestrator_Orchestrate_Format(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		format     string
		expected   string
	}{
		{"default", "Invalid", "default", "Error: Validation error: The specified value \"Invalid\" is invalid. Issues: must be in lower case"},
		{"github-actions", "Invalid", "github-actions", "::error::Validation error: The specified value \"Invalid\" is invalid. Issues: must be in lower case"},
	}

	for _, tc := range cases {
		sut := &Orchestrator{
			Value:     &Value{raw: tc.value},
			Validator: &Validator{Errors: &Errors{}, lowerCase: true},
			Formatter: &Formatter{format: tc.format},
		}
		err := sut.Orchestrate()

		format := "\n annotation: %s\n expected:   %s\n actual:     %+v\n value:      %s\n format:     %v"
		if err == nil {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, NoError, tc.value, tc.format))
		} else if err.Error() != tc.expected {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, err.Error(), tc.value, tc.format))
		}
	}
}
