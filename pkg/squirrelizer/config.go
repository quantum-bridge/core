package squirrelizer

import (
	"database/sql"
	"github.com/pkg/errors"
	"sync"
	"time"
)

// Database is the interface for the database client.
type Database interface {
	// Instance returns the database client instance.
	Instance() *DB
	// SQLInstance returns the SQL database client instance.
	SQLInstance() *sql.DB
}

// DBConfig is the structure that holds the database configuration.
type DBConfig struct {
	URL                      string        `config:"url,required"`
	MaxOpenConnections       int           `config:"max_open_connections"`
	MaxIdleConnections       int           `config:"max_idle_connections"`
	MinListenerRetryDuration time.Duration `config:"min_listener_retry_duration"`
	MaxListenerRetryDuration time.Duration `config:"max_listener_retry_duration"`
}

// database is the structure that holds the database client data and configuration.
type database struct {
	config DBConfig
	once   sync.Once
}

// NewDatabase creates a new database client instance with pre-loaded configuration.
func NewDatabase(config DBConfig) Database {
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
