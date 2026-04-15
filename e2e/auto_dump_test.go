package e2e_tests

import (
	"os"
	"testing"
    "github.com/guilhermewebdev/migrator/src/cli"
)

func TestAutoDumpSchema(t *testing.T) {
    os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/e2e/mocks/1")
	env(func(envs env_set) {
		schema_file := os.Getenv("ROOT_DIR") + "/tmp/auto_schema.sql"
		os.Remove(schema_file)

		os.Setenv("AUTO_DUMP_SCHEMA", "true")
		os.Setenv("SCHEMA_FILE_PATH", schema_file)

		_, err := capture_output(func() error {
			return cli.Run([]string{"migrator", "up"})
		})
		if err != nil {
			t.Fatalf("Failed to run up for %s: %v", envs["DB_DRIVER"], err)
		}

		if !file_exists(schema_file) {
			t.Errorf("Auto dump schema file was not created for %s", envs["DB_DRIVER"])
		}
        
        // Clean up for next env
        os.Remove(schema_file)
	})
}
