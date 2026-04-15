package e2e_tests

import (
	"os"
	"testing"
    "github.com/guilhermewebdev/migrator/src/cli"
)

func TestDumpSchema(t *testing.T) {
	env(func(envs env_set) {
		schema_file := os.Getenv("ROOT_DIR") + "/tmp/schema.sql"
		os.Remove(schema_file)

		os.Setenv("SCHEMA_FILE_PATH", schema_file)
		_, err := capture_output(func() error {
			return cli.Run([]string{"migrator", "schema"})
		})
		if err != nil {
			t.Fatalf("Failed to dump schema for %s: %v", envs["DB_DRIVER"], err)
		}

		if !file_exists(schema_file) {
			t.Errorf("Schema file was not created for %s", envs["DB_DRIVER"])
		}

		content, err := os.ReadFile(schema_file)
		if err != nil {
			t.Fatalf("Could not read schema file: %v", err)
		}

		if len(content) == 0 {
			t.Errorf("Schema file is empty for %s", envs["DB_DRIVER"])
		}
	})
}
