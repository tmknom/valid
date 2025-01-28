package internal

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
)

func TestApp_Run_Valid(t *testing.T) {
	cases := []struct {
		annotation string
		args       []string
	}{
		{
			annotation: "number",
			args:       []string{"--exact-length", "5", "--digit", "--value", "12345"},
		},
	}

	for _, tc := range cases {
		sut := NewApp(FakeTestIO())
		err := sut.Run(context.Background(), tc.args)

		format := "\n expected: %s\n actual:   %s\n args:     %v"
		if err != nil {
			t.Errorf(fmt.Sprintf(format, NoError, err, tc.args))
		}
	}
}

func TestApp_Run_Invalid(t *testing.T) {
	cases := []struct {
		annotation string
		args       []string
		expected   string
	}{
		{
			annotation: "validation_error",
			args:       []string{"--exact-length", "5", "--digit", "--value", "123a"},
			expected:   "Error: Validation error: The specified value \"123a\" is invalid. Issues: the length must be exactly 5, must contain digits only.",
		},
		{
			annotation: "argument_error",
			args:       []string{"--min", "5", "--exact-length", "abc", "--alphanumeric", "--value", "123a"},
			expected:   "Error: Argument error: --min cannot validate non-numeric value, --exact-length must be an integer number.",
		},
		{
			annotation: "complex_error",
			args:       []string{"--value-name", "test-id", "--mask-value", "--format", "github-actions", "--exact-length", "abc", "--digit", "--upper-case", "--value", "123a"},
			expected:   "::error::Validation error: The specified test-id \"***\" is invalid. Issues: must contain digits only, must be in upper case; Argument error: --exact-length must be an integer number.",
		},
	}

	for _, tc := range cases {
		sut := NewApp(FakeTestIO())
		err := sut.Run(context.Background(), tc.args)

		format := "\n expected: %s\n actual:   %s\n args:     %v"
		if err == nil {
			t.Errorf(fmt.Sprintf(format, tc.expected, NoError, tc.args))
		} else if err.Error() != tc.expected {
			t.Errorf(fmt.Sprintf(format, tc.expected, err, tc.args))
		}
	}
}

func FakeTestIO() *IO {
	return &IO{
		InReader:  &bytes.Buffer{},
		OutWriter: &bytes.Buffer{},
		ErrWriter: os.Stderr,
	}
}
