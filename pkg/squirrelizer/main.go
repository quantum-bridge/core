package squirrelizer

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Config is the configuration for the database client.
type Config struct {
	// URL is the connection URL for the database.
	URL string
	// MaxOpenConnections is the maximum number of open connections to the database.
	MaxOpenConnections int
	// MaxIdleConnections is the maximum number of idle connections to the database.
	MaxIdleConnections int
}

// Open opens a new database connection with the provided configuration.
func Open(cfg Config) (*DB, error) {
	// Connect to the database using the provided connection URL.
	db, err := sqlx.Connect("postgres", cfg.URL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	// Set the maximum number of open and idle connections.
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.MaxOpenConnections)

	// Return the database client instance.
	return &DB{
		sqlxDB: db,
		Query:  newQueryHandler(db),
	}, nil
}

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

// TransactionHandler is the function that is being used in the transaction handler.
type TransactionHandler func() error

// Transactional is the interface for the transaction methods.
type Transactional interface {
	// ExecuteTransaction executes the transaction function without any options.
	ExecuteTransaction(handler TransactionHandler) error
}

// Connection is the interface for the database connection.
type Connection interface {
	// Transactional is the interface for the transaction methods.
	Transactional
	// Query is the interface for the database query methods.
	Query
}

// Query is the interface for the database query methods.
type Query interface {
	// Executor is the interface for the SQL executor methods.
	Executor
	// Selector is the interface for the SQL selector methods.
	Selector
	// Getter is the interface for the SQL getter methods.
	Getter
}
