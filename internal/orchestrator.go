package internal

func newOrchestrator() *Orchestrator {
	return &Orchestrator{
		Value:     &Value{},
		Validator: &Validator{Errors: &Errors{}},
	}
}

type Orchestrator struct {
	*Value
	*Validator
}

func (o *Orchestrator) orchestrate() error {
	o.Validator.UnmaskedValue = o.Value.Unmasked()
	o.Validator.Errors.MaskedValue = o.Value.Masked()
	return o.Validator.validate()
}

type Value struct {
	raw  string
	mask bool
}

func (v *Value) Unmasked() string {
	return v.raw
}

func (v *Value) Masked() string {
	if v.mask {
		return MaskedValue
	}
	return v.raw
}

const MaskedValue = "***"
