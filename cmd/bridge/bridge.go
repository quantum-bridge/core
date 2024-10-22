package bridge

import (
	"github.com/alecthomas/kingpin"
	"github.com/quantum-bridge/core/cmd/config"
	"github.com/quantum-bridge/core/cmd/env"
	"github.com/quantum-bridge/core/cmd/service"
	"go.uber.org/zap"
)

// Run starts the bridge service with the given configuration.
func Run(args []string) {
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

	// Application instance with the name and description.
	application := kingpin.New("bridge", "Core bridge API")

	// Commands for the running of the service.
	runCommand := application.Command("run", "run the bridge service")
	runService := runCommand.Command("service", "run the bridge service")

	// Commands for the migration of the database.
	migrateCommand := application.Command("migrate", "migrate up the database")
	migrateUpCommand := migrateCommand.Command("up", "migrate up the database")
	migrateDownCommand := migrateCommand.Command("down", "migrate down the database")

	parameters, err := application.Parse(args[1:3])
	if err != nil {
		logger.Error("error parsing arguments", zap.Error(err))

		return
	}

	switch parameters {
	case runService.FullCommand():
		servicesMap := make(map[string]bool)
		for _, svc := range args[3:] {
			servicesMap[svc] = true
		}

		cfg.SetServicesConfig(&config.ServicesConfig{
			TxHistory: servicesMap["tx-history"],
		})

		// Start the Core BackEnd service.
		service.Run(cfg, logger)
	case migrateUpCommand.FullCommand():
		// Migrate up the database.
		MigrateUp(cfg, logger)
	case migrateDownCommand.FullCommand():
		// Migrate down the database.
		MigrateDown(cfg, logger)
	default:
		logger.Errorf("unknown command: %s", parameters)
	}
}
