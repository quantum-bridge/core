package configure

import (
	"time"

	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/env"
	"github.com/quantum-bridge/core/pkg/squirrelizer"
)

// DB returns the database client that is being used in the bridge.
func (c *config) DB() *squirrelizer.DB {
	var db *squirrelizer.DB

	c.onceDB.Do(func() {
		cfg := struct {
			URL                      string        `config:"url,required"`
			MaxOpenConnections       int           `config:"max_open_connections"`
			MaxIdleConnections       int           `config:"max_idle_connections"`
			MinListenerRetryDuration time.Duration `config:"min_listener_retry_duration"`
			MaxListenerRetryDuration time.Duration `config:"max_listener_retry_duration"`
		}{}

		// Load the database configuration.
		configData, err := c.getter.GetStringMap("db")
		if err != nil {
			panic(errors.Wrap(err, "failed to retrieve database configuration"))
		}

		// Load the database configuration.
		if err := env.NewConfiguration(&cfg).From(configData).Load(); err != nil {
			panic(errors.Wrap(err, "failed to load database configuration"))
		}

		// Initialize the database client with the loaded configuration.
		db = squirrelizer.NewDatabase(cfg).Instance()
	})

	return db
}
