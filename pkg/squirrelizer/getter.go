package squirrelizer

import (
	"context"
	"github.com/Masterminds/squirrel"
)

// Getter is the interface for the SQL getter methods.
type Getter interface {
	// Get executes the query and maps the result to the provided destination without any context.
	Get(dest interface{}, query squirrel.Sqlizer) error
	// GetContext executes the query and maps the result to the provided destination with context.
	GetContext(ctx context.Context, dest interface{}, query squirrel.Sqlizer) error
	// GetRaw executes the query and maps the result to the provided destination without any context.
	GetRaw(dest interface{}, query string, args ...interface{}) error
	// GetRawContext executes the query and maps the result to the provided destination with context.
	GetRawContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
