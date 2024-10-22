package squirrelizer

import (
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
