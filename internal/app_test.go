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
			args:       []string{"--exactly-length", "5", "--digit", "--value", "12345"},
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
			args:       []string{"--exactly-length", "5", "--digit", "--value", "123a"},
			expected:   "the length must be exactly 5, must contain digits only",
		},
		{
			annotation: "argument_error",
			args:       []string{"--exactly-length", "abc", "--digit", "--value", "123"},
			expected:   "strconv.Atoi: parsing \"abc\": invalid syntax",
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
	return fmt.Sprintf("expected: %s, actual: %#v, args: %v", expected, actual, args)
}

func FakeTestIO() *IO {
	return &IO{
		InReader:  &bytes.Buffer{},
		OutWriter: &bytes.Buffer{},
		ErrWriter: os.Stderr,
	}
}
