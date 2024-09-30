package squirrelizer

import (
	"context"
	"github.com/Masterminds/squirrel"
)

// Selector is the interface for the SQL selector methods.
type Selector interface {
	// Select executes the query and maps the result to the provided destination without any context.
	Select(dest interface{}, query squirrel.Sqlizer) error
	// SelectContext executes the query and maps the result to the provided destination with context.
	SelectContext(ctx context.Context, dest interface{}, query squirrel.Sqlizer) error
	// SelectRaw executes the query and maps the result to the provided destination without any context.
	SelectRaw(dest interface{}, query string, args ...interface{}) error
	// SelectRawContext executes the query and maps the result to the provided destination with context.
	SelectRawContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
