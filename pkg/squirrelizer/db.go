package squirrelizer

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// DB is the database client that is being used in the bridge.
type DB struct {
	Query
	sqlxDB *sqlx.DB
}

// SQLInstance returns the SQL database client instance.
func (db *DB) SQLInstance() *sql.DB {
	return db.sqlxDB.DB
}

// Clone returns a new database client instance.
func (db *DB) Clone() *DB {
	return &DB{
		Query:  newQueryHandler(db.sqlxDB),
		sqlxDB: db.sqlxDB,
	}
}

// TransactionFunc is the function that is being used in the transaction handler.
type TransactionFunc func() error

// ExecuteTransaction executes the transaction function without any options.
func (db *DB) ExecuteTransaction(fn TransactionFunc) error {
	return db.ExecuteTransactionWithOptions(nil, fn)
}

// ExecuteTransactionWithOptions executes the transaction function with options.
func (db *DB) ExecuteTransactionWithOptions(opts *sql.TxOptions, fn TransactionFunc) error {
	// Start the transaction using the provided transaction options.
	tx, err := db.sqlxDB.BeginTxx(context.Background(), opts)
	if err != nil {
		// If there's an error starting the transaction, wrap and return the error.
		return errors.Wrap(err, "failed to start transaction")
	}

	// Set the Query to a new query handler with the transaction.
	db.Query = newQueryHandler(tx)

	// Ensure the transaction is rolled back if any error occurs. A
	defer tx.Rollback()

	// Reset the Query to a new query handler with the original database connection after the transaction is done.
	defer func() {
		db.Query = newQueryHandler(db.sqlxDB)
	}()

	// Execute the provided transaction function.
	if err := fn(); err != nil {
		// If there's an error executing the transaction function, wrap and return the error.
		return errors.Wrap(err, "transaction failed")
	}

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "transaction commit failed")
	}

	return nil
}
