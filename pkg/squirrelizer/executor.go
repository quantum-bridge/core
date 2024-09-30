package squirrelizer

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
)

// Executor is the interface for the SQL executor methods.
type Executor interface {
	// Exec executes the query without any context.
	Exec(query squirrel.Sqlizer) error
	// ExecContext executes the query with context.
	ExecContext(ctx context.Context, query squirrel.Sqlizer) error
	// ExecRaw executes the query without any context.
	ExecRaw(query string, args ...interface{}) error
	// ExecRawContext executes the query with context.
	ExecRawContext(ctx context.Context, query string, args ...interface{}) error
	// ExecWithResult executes the query and returns the result without any context.
	ExecWithResult(query squirrel.Sqlizer) (sql.Result, error)
	// ExecWithResultContext executes the query and returns the result with context.
	ExecWithResultContext(ctx context.Context, query squirrel.Sqlizer) (sql.Result, error)
}
