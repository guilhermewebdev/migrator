package e2e_tests

import (
	"os"
	"testing"

	"github.com/guilhermewebdev/migrator/src/cli"
)

func TestDown(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/e2e/mocks/1")
	env(func(envs env_set) {
		t.Log("\nTesting with envs: ", envs, "\n")
		if err := cli.Run([]string{"migrator", "up"}); err != nil {
			t.Fatal(err)
		}
		if err := cli.Run([]string{"migrator", "down"}); err != nil {
			t.Fatal(err)
		}
	})
}

func TestDown_WithoutMigrations(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/e2e/mocks/1")
	env(func(envs env_set) {
		t.Log("\nTesting with envs: ", envs, "\n")
		err := cli.Run([]string{"migrator", "down"})
		if err == nil || err.Error() != "No migrations to rollback" {
			t.Fatal(err)
		}
	})
}
