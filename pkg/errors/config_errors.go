package errors

import "github.com/pkg/errors"

var (
	// ErrNonZeroValue is an error that is returned when a field is required to be non-zero.
	ErrNonZeroValue = errors.New("you must set non zero value to this field")
	// ErrRequiredValue is an error that is returned when a field is required to be set.
	ErrRequiredValue = errors.New("you must set the value in field")
	// ErrNotValid is an error that is returned when a value is not valid.
	ErrNotValid = errors.New("not valid value")
	// ErrNoBackends is an error that is returned when no backends are configured.
	ErrNoBackends = errors.New("no backends configured")
)
