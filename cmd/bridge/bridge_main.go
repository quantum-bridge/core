package bridge

import (
	"github.com/quantum-bridge/core/cmd/config"
	"github.com/quantum-bridge/core/cmd/env"
	"github.com/quantum-bridge/core/cmd/service"
	"go.uber.org/zap"
)

// Run starts the bridge service with the given configuration.
func Run() {
	// Create a new logger.
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	// Create a new logger instance.
	logger := log.Sugar()

	defer func() {
		// Recover from panics and log them.
		if rvr := recover(); rvr != nil {
			logger.Error("app panicked", zap.Any("recover", rvr))
		}
	}()

	// Load the configuration from the environment.
	cfg := config.New(env.MustFromEnv())

	// Start the BackEnd service.
	service.Run(cfg, logger)
}
