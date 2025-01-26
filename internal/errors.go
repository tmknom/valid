package internal

import (
	"fmt"
	"strings"
)

type Errors struct {
	*Value
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
	if e.hasValidations() && e.hasArguments() {
		return strings.Join([]string{e.joinValidationError(), e.joinArgumentError()}, "; ")
	} else if e.hasValidations() {
		return e.joinValidationError()
	} else if e.hasArguments() {
		return e.joinArgumentError()
	} else {
		return ""
	}
}

func (e *Errors) joinValidationError() string {
	if !e.hasValidations() {
		return ""
	}

	issues := make([]string, 0, len(e.validations))
	for _, err := range e.validations {
		issues = append(issues, err.Error())
	}
	message := fmt.Sprintf("Validation error: The value \"%s\" is invalid. Issues: %s", e.Value.Masked(), strings.Join(issues, ", "))
	return message
}

func (e *Errors) joinArgumentError() string {
	if !e.hasArguments() {
		return ""
	}

	issues := make([]string, 0, len(e.arguments))
	for _, err := range e.arguments {
		issues = append(issues, err.Error())
	}
	return "Argument error: " + strings.Join(issues, ", ")
}
