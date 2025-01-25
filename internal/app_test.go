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

		if err != nil {
			t.Errorf(messageWithArgs(NoError, err, tc.args))
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
			expected:   "Validation error: The value \"123a\" is invalid. Issues: the length must be exactly 5, must contain digits only",
		},
		{
			annotation: "argument_error",
			args:       []string{"--min", "5", "--exact-length", "abc", "--alphanumeric", "--value", "123a"},
			expected:   "Argument error: --min cannot validate \"123a\", --exact-length must be an integer number",
		},
		{
			annotation: "validation_and_argument_error",
			args:       []string{"--exact-length", "abc", "--digit", "--value", "123a"},
			expected:   "Validation error: The value \"123a\" is invalid. Issues: must contain digits only; Argument error: --exact-length must be an integer number",
		},
	}

	for _, tc := range cases {
		sut := NewApp(FakeTestIO())
		err := sut.Run(context.Background(), tc.args)

		if err == nil {
			t.Errorf(messageWithArgs(tc.expected, NoError, tc.args))
		} else if err.Error() != tc.expected {
			t.Errorf(messageWithArgs(tc.expected, err, tc.args))
		}
	}
}

func messageWithArgs(expected string, actual any, args []string) string {
	return fmt.Sprintf("\n expected: %s\n actual:   %s\n args:     %v", expected, actual, args)
}

func FakeTestIO() *IO {
	return &IO{
		InReader:  &bytes.Buffer{},
		OutWriter: &bytes.Buffer{},
		ErrWriter: os.Stderr,
	}
}
