package errs

import "fmt"

var configurationErrPrefix = getErrPrefix("Configuration")

type ConfigurationError struct {
	Details string
	Err     error
}

func (e *ConfigurationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %s", configurationErrPrefix, e.Details, e.Err.Error())
	}
	return fmt.Sprintf("%s: %s", configurationErrPrefix, e.Details)
}

func NewConfigurationErr(opt Opts) *ConfigurationError {
	return &ConfigurationError{
		Details: opt.Details,
		Err:     opt.Error,
	}
}
