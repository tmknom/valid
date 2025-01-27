package internal

import "fmt"

type Formatter struct {
	format string
}

func (f *Formatter) Format(err error) error {
	if err == nil {
		return nil
	}
	switch f.format {
	case "github-actions":
		return fmt.Errorf("::error::%s", err.Error())
	default:
		return fmt.Errorf("Error: %s", err.Error())
	}
}
