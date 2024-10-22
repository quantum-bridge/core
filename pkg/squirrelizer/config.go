package squirrelizer

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/pkg/squirrelizer/shared"
	"sync"
)

// Database is the interface for the database client.
type Database interface {
	// Instance returns the database client instance.
	Instance() *DB
	// SQLInstance returns the SQL database client instance.
	SQLInstance() *sql.DB
}

// database is the structure that holds the database client data and configuration.
type database struct {
	config shared.DBConfig
	once   sync.Once
}

// NewDatabase creates a new database client instance with preloaded configuration.
func NewDatabase(config shared.DBConfig) Database {
	return &database{
		config: config,
	}
}

// Instance returns the database client instance.
func (d *database) Instance() *DB {
	var db *DB
	var err error

	d.once.Do(func() {
		db, err = Open(Config{
			URL:                d.config.URL,
			MaxOpenConnections: d.config.MaxOpenConnections,
			MaxIdleConnections: d.config.MaxIdleConnections,
		})
		if err != nil {
			panic(errors.Wrap(err, "failed to open database connection"))
		}
	})

	return db
}

// SQLInstance returns the SQL database client instance.
func (d *database) SQLInstance() *sql.DB {
	return d.Instance().SQLInstance()
}
