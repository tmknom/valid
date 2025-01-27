package internal

import (
	"fmt"
	"testing"
)

func TestValue_Name(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		name       string
		expected   string
	}{
		{"specified", "test-value", "test-id", "test-id"},
		{"not-specified", "test-value", "", "value"},
	}

	for _, tc := range cases {
		sut := &Value{raw: tc.value, name: tc.name}
		actual := sut.Name()

		format := "\n annotation: %s\n expected:   %s\n actual:     %+v\n value:      %s\n name:       %s"
		if tc.expected != actual {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, actual, tc.value, tc.name))
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
