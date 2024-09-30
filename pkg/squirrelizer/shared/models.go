package shared

import "time"

// DBConfig is the structure that holds the database configuration.
type DBConfig struct {
	URL                      string        `config:"url,required"`
	MaxOpenConnections       int           `config:"max_open_connections"`
	MaxIdleConnections       int           `config:"max_idle_connections"`
	MinListenerRetryDuration time.Duration `config:"min_listener_retry_duration"`
	MaxListenerRetryDuration time.Duration `config:"max_listener_retry_duration"`
}
