package bridge

import (
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/config"
	sqlmigration "github.com/quantum-bridge/core/cmd/migrations"
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
)

// Migration source for embedded files
var migrations = &migrate.EmbedFileSystemMigrationSource{
	FileSystem: sqlmigration.Migrations,
	Root:       "sql", // Directory in the embedded file system where your migrations are stored
}

// MigrateUp applies new migrations to the database
func MigrateUp(cfg config.Config, logger *zap.SugaredLogger) {
	n, err := migrate.Exec(cfg.DB().SQLInstance(), "postgres", migrations, migrate.Up)
	if err != nil {
		panic(errors.Wrap(err, "Migration failed"))
	}

	logger.Logf(zap.InfoLevel, "Applied %d migrations!", n)
}

// MigrateDown rolls back the latest migrations
func MigrateDown(cfg config.Config, logger *zap.SugaredLogger) {
	n, err := migrate.Exec(cfg.DB().SQLInstance(), "postgres", migrations, migrate.Down)
	if err != nil {
		panic(errors.Wrap(err, "Migration failed"))
	}

	logger.Logf(zap.InfoLevel, "Rolled back %d migrations!", n)
}
