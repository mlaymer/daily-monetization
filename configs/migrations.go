package configs

import "github.com/caarlos0/env/v6"

// Migrations contains configuration parameters
// for performing migrations at application start.
type Migrations struct {
	// SourceURL defines the URL of the source
	// of files with migrations. To specify the
	// directory with migration files in the file
	// system, use the syntax file://path_to_migrates.
	//
	// The default is file://migrations.
	SourceURL string `env:"MIGRATIONS_SOURCE_URL" envDefault:"file://migrations"`

	// Version is used to tell the migration tool
	// to which version to apply up or down migrations.
	//
	// This parameter is required.
	Version uint `env:"MIGRATIONS_VERSION,required"`

	// Driver contains the name of the driver that
	// will be used when connecting and working with MySQL.
	//
	// By default, the "mysql" driver is used, you do not
	// need to change the driver unless you really need to.
	//
	// Alternative driver list: https://golang.org/s/sqldrivers.
	Driver string `env:"MIGRATIONS_DRIVER" envDefault:"mysql"`

	// The DSN contains a data source name that contains a specific
	// string for connecting and working with MySQL.
	//
	// This parameter is required.
	DSN string `env:"MIGRATIONS_DSN,required"`

	// The database contains the name of the database
	// to perform migrations.
	//
	// The default is "mysql".
	Database string `env:"MIGRATIONS_DATABASE" envDefault:"mysql"`
}

// LoadMigrations loads configuration parameters for
// performing migrations from environment variables.
//
// If an error occurs while loading the configuration
// parameters, the method will return it.
func LoadMigrations() (*Migrations, error) {
	config := new(Migrations)

	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}
