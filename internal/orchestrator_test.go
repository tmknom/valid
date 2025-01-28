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
		{"invalid-masked", "InvalidMasked", true, "Error: Validation error: The specified value \"***\" is invalid. Issues: must be in lower case"},
		{"invalid-unmasked", "InvalidUnmasked", false, "Error: Validation error: The specified value \"InvalidUnmasked\" is invalid. Issues: must be in lower case"},
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
