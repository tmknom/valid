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

func (o *Orchestrator) orchestrate() error {
	o.Validator.UnmaskedValue = o.Value.Unmasked()
	o.Validator.Errors.value = o.Value
	return o.Formatter.Format(o.Validator.validate())
}

type Value struct {
	raw  string
	name string
	mask bool
}

func (v *Value) Name() string {
	if len(v.name) == 0 {
		return DefaultValueName
	}
	return v.name
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

const DefaultValueName = "value"
const MaskedValue = "***"
