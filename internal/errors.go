package internal

import (
	"strings"
)

type Errors struct {
	validations []error
	arguments   []error
}

func (e *Errors) AddValidationError(err error) {
	e.validations = append(e.validations, err)
}

func (e *Errors) AddArgumentError(err error) {
	e.arguments = append(e.arguments, err)
}

func (e *Errors) HasError() bool {
	return e.hasValidations() || e.hasArguments()
}

func (e *Errors) hasValidations() bool {
	return len(e.validations) > 0
}

func (e *Errors) hasArguments() bool {
	return len(e.arguments) > 0
}

func (e *Errors) Error() string {
	messages := make([]string, 0, len(e.validations)+len(e.arguments))
	for _, err := range e.validations {
		messages = append(messages, err.Error())
	}
	for _, err := range e.arguments {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, ", ")
}
