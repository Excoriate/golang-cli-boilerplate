package errs

import "fmt"

var commandErrAPrefix = getErrPrefix("Command")

type CommandError struct {
	Details string
	Err     error
}

func (e *CommandError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %s", commandErrAPrefix, e.Details, e.Err.Error())
	}
	return fmt.Sprintf("%s: %s", commandErrAPrefix, e.Details)
}

func NewCommandErr(opt Opts) *CommandError {
	return &CommandError{
		Details: opt.Details,
		Err:     opt.Error,
	}
}
