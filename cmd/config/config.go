package config

import (
	"github.com/quantum-bridge/core/cmd/config/configure"
	"github.com/quantum-bridge/core/cmd/env"
)

// Config is the interface for the configuration of the bridge.
type Config interface {
	configure.Config
}

// config is struct that holds the configuration of the bridge.
type config struct {
	configure.Config
	getter env.Getter
}

// New creates a new configuration instance with the given getter.
func New(getter env.Getter) Config {
	return &config{
		getter: getter,
		Config: configure.NewConfig(getter),
	}
}
