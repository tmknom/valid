package internal

func newOrchestrator() *Orchestrator {
	return &Orchestrator{
		Value:     &Value{},
		Validator: &Validator{Errors: &Errors{}},
		Formatter: &Formatter{},
	}
}

type Orchestrator struct {
	*Value
	*Validator
	*Formatter
}

func (o *Orchestrator) Orchestrate() error {
	o.Validator.UnmaskedValue = o.Value.Unmasked()
	o.Validator.Errors.value = o.Value
	return o.Formatter.Format(o.Validator.Validate())
}
