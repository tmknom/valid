package internal

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
