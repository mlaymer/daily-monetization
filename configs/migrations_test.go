package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadMigrations(t *testing.T) {
	t.Parallel()

	// Arrange.
	tests := []struct {
		name      string
		arrange   func()
		assert    func(config *Migrations)
		willBeErr bool
	}{
		{
			name:    "Check the default",
			arrange: func() {},
			assert: func(config *Migrations) {
				assert.Equal(t, "file://migrations", config.SourceURL)
			},
		},
		{
			name: "Check the load from environment variables",
			arrange: func() {
				assert.NoError(t, os.Setenv("MIGRATIONS_SOURCE_URL", "file://monetization/migrations"))
			},
			assert: func(config *Migrations) {
				assert.Equal(t, "file://monetization/migrations", config.SourceURL)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.arrange()

			// Act.
			actualConfig, actualErr := LoadMigrations()

			// Assert.
			require.NoError(t, actualErr)

			test.assert(actualConfig)
		})
	}
}
