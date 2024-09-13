package errors

import "github.com/pkg/errors"

var (
	// ErrNotFound is an error that is returned when a requested item is not found.
	ErrNotFound = errors.New("not found")
	// ErrTxNotFound is an error that is returned when a transaction is not found.
	ErrTxNotFound = errors.New("transaction not found")
	// ErrTxFailed is an error that is returned when a transaction failed.
	ErrTxFailed = errors.New("transaction failed")
	// ErrTxNotConfirmed is an error that is returned when a transaction is not confirmed.
	ErrTxNotConfirmed = errors.New("transaction not confirmed")
	// ErrEventNotFound is an error that is returned when an event is not found.
	ErrEventNotFound = errors.New("event not found")
	// ErrWrongLockEvent is an error that is returned when a lock event is wrong.
	ErrWrongLockEvent = errors.New("wrong lock event")
	// ErrAlreadyWithdrawn is an error that is returned when a withdrawal is already done.
	ErrAlreadyWithdrawn = errors.New("already withdrawn")
)
