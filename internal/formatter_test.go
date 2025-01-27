package internal

import (
	"fmt"
	"testing"
)

func TestFormatter_format(t *testing.T) {
	cases := []struct {
		annotation string
		value      string
		expected   string
	}{
		{"default", "default", "Error: example error"},
		{"github-actions", "github-actions", "::error::example error"},
		{"invalid", "invalid", "Error: example error"},
	}

	for _, tc := range cases {
		sut := &Formatter{
			format: tc.value,
		}
		err := sut.Format(fmt.Errorf("example error"))

		format := "\n annotation: %s\n expected:   %s\n actual:     %+v\n value:      %s"
		if err == nil {
			t.Fatalf(fmt.Sprintf(format, tc.annotation, tc.expected, NoError, tc.value))
		}
		if err.Error() != tc.expected {
			t.Errorf(fmt.Sprintf(format, tc.annotation, tc.expected, err, tc.value))
		}
	}
}
