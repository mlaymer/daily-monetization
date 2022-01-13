package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/dailydotdev/daily-monetization/configs"
)

// Migrations provides an API for high-level
// management of application migrations.
//
// Migrations connect to the database, executes
// migration requests and releases resources after
// completion.
//
// Migrations use MigrationsLogger for logging
// during migrations.
type Migrations struct {
	config *configs.Migrations

	database *sql.DB
	migrates *migrate.Migrate
}

func (api *Migrations) Close() error {
	if err := api.database.Close(); err != nil {
		return err
	}

	sourceErr, dbErr := api.migrates.Close()
	if sourceErr != nil {
		return sourceErr
	}

	if dbErr != nil {
		return dbErr
	}

	return nil
}

// Apply looks at the currently active migration
// version and then migrates up or down to the
// version specified in the configuration settings.
//
// See configs.Migrations.
func (api *Migrations) Apply() error {
	return api.migrates.Migrate(api.config.Version)
}

// Drop deletes everything in the database.
func (api *Migrations) Drop() error {
	return api.migrates.Drop()
}

// NewMigrations initializes the service to apply
// migrations to the database.
//
// The method connects to the database and prepares
// the migrations to be applied.
func NewMigrations(config *configs.Migrations) (*Migrations, error) {
	connection, err := sql.Open(config.Driver, config.DSN)
	if err != nil {
		return nil, err
	}

	if err := connection.Ping(); err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(connection, new(mysql.Config))
	if err != nil {
		return nil, err
	}

	instance, err := migrate.NewWithDatabaseInstance(config.SourceURL, config.Database, driver)
	if err != nil {
		return nil, err
	}

	return &Migrations{
		config:   config,
		database: connection,
		migrates: instance,
	}, nil
}
