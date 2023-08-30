package errs

import "fmt"

var adapterErrorPrefix = getErrPrefix("Adapter")

type AdapterError struct {
	Details string
	Err     error
}

func (e *AdapterError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %s", adapterErrorPrefix, e.Details, e.Err.Error())
	}
	return fmt.Sprintf("%s: %s", adapterErrorPrefix, e.Details)
}

func NewAdapterErr(opt Opts) *AdapterError {
	return &AdapterError{
		Details: opt.Details,
		Err:     opt.Error,
	}
}
