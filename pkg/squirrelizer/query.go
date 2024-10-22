package squirrelizer

import (
	"context"
	"database/sql"
	"github.com/quantum-bridge/core/pkg/squirrelizer/shared"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Query is the interface for the database query methods.
type Query interface {
	// Executor is the interface for the SQL executor methods.
	Executor
	// Selector is the interface for the SQL selector methods.
	Selector
	// Getter is the interface for the SQL getter methods.
	Getter
}

// sqlExecutor is the interface for the SQL executor.
type sqlExecutor interface {
	// ExecerContext is the interface for the SQL executor context methods.
	sqlx.ExecerContext
	// SelectContext executes the query and maps the result to the provided destination with context.
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// GetContext executes the query and maps the result to the provided destination with context.
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// queryHandler is the structure that holds the executor SQL executor.
type queryHandler struct {
	executor sqlExecutor
}

// newQueryHandler creates a new query handler instance with the provided raw SQL executor.
func newQueryHandler(executor sqlExecutor) *queryHandler {
	return &queryHandler{
		executor: executor,
	}
}

// Get executes the query and maps the result to the provided destination without any context.
func (q *queryHandler) Get(dest interface{}, query squirrel.Sqlizer) error {
	return q.GetContext(context.Background(), dest, query)
}

// GetContext executes the query and maps the result to the provided destination with context.
func (q *queryHandler) GetContext(ctx context.Context, dest interface{}, query squirrel.Sqlizer) error {
	sqlStatement, args, err := shared.Build(query)
	if err != nil {
		return err
	}

	return q.GetRawContext(ctx, dest, sqlStatement, args...)
}

// GetRaw executes the query and maps the result to the provided destination without any context.
func (q *queryHandler) GetRaw(dest interface{}, query string, args ...interface{}) error {
	return q.GetRawContext(context.Background(), dest, query, args...)
}

// GetRawContext executes the query and maps the result to the provided destination with context.
func (q *queryHandler) GetRawContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	query = shared.Rebind(query)
	err := q.executor.GetContext(ctx, dest, query, args...)

	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return errors.Wrap(err, "failed to get raw")
}

// Exec executes the query without any context.
func (q *queryHandler) Exec(query squirrel.Sqlizer) error {
	return q.ExecContext(context.Background(), query)
}

// ExecContext executes the query with context.
func (q *queryHandler) ExecContext(ctx context.Context, query squirrel.Sqlizer) error {
	sqlStatement, args, err := shared.Build(query)
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}
	return q.ExecRawContext(ctx, sqlStatement, args...)
}

// ExecRaw executes the query without any context.
func (q *queryHandler) ExecRaw(query string, args ...interface{}) error {
	return q.ExecRawContext(context.Background(), query, args...)
}

// ExecRawContext executes the query with context.
func (q *queryHandler) ExecRawContext(ctx context.Context, query string, args ...interface{}) error {
	query = shared.Rebind(query)
	_, err := q.executor.ExecContext(ctx, query, args...)
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return errors.Wrap(err, "failed to exec query")
}

// ExecWithResult executes the query and returns the result without any context.
func (q *queryHandler) ExecWithResult(query squirrel.Sqlizer) (sql.Result, error) {
	return q.ExecWithResultContext(context.Background(), query)
}

// ExecWithResultContext executes the query and returns the result with context.
func (q *queryHandler) ExecWithResultContext(ctx context.Context, query squirrel.Sqlizer) (sql.Result, error) {
	sqlStatement, args, err := shared.Build(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}
	return q.ExecRawWithResultContext(ctx, sqlStatement, args...)
}

// ExecRawWithResult executes the query and returns the result without any context.
func (q *queryHandler) ExecRawWithResult(query string, args ...interface{}) (sql.Result, error) {
	return q.ExecRawWithResultContext(context.Background(), query, args...)
}

// ExecRawWithResultContext executes the query and returns the result with context.
func (q *queryHandler) ExecRawWithResultContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	query = shared.Rebind(query)
	result, err := q.executor.ExecContext(ctx, query, args...)
	if err == nil {
		return result, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return nil, errors.Wrap(err, "failed to exec query")
}

// Select executes the query and maps the result to the provided destination without any context.
func (q *queryHandler) Select(dest interface{}, query squirrel.Sqlizer) error {
	return q.SelectContext(context.Background(), dest, query)
}

// SelectContext executes the query and maps the result to the provided destination with context.
func (q *queryHandler) SelectContext(ctx context.Context, dest interface{}, query squirrel.Sqlizer) error {
	sqlStatement, args, err := shared.Build(query)
	if err != nil {
		return err
	}
	return q.SelectRawContext(ctx, dest, sqlStatement, args...)
}

// SelectRaw executes the query and maps the result to the provided destination without any context.
func (q *queryHandler) SelectRaw(dest interface{}, query string, args ...interface{}) error {
	return q.SelectRawContext(context.Background(), dest, query, args...)
}

// SelectRawContext executes the query and maps the result to the provided destination with context.
func (q *queryHandler) SelectRawContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	query = shared.Rebind(query)
	err := q.executor.SelectContext(ctx, dest, query, args...)

	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return errors.Wrap(err, "failed to select")
}
