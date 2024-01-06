package e2e_tests

import (
	"os"
	"testing"

	"github.com/guilhermewebdev/migrator/src/cli"
)

func TestLatest(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/e2e/mocks/1")
	env(func(envs env_set) {
		t.Log("\nTesting with envs: ", envs, "\n")
		if err := cli.Run([]string{"migrator", "latest"}); err != nil {
			t.Fatal(err)
		}
	})
}

func TestLatest_WhenMigrationsWereRan(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/e2e/mocks/1")
	env(func(envs env_set) {
		t.Log("\nTesting with envs: ", envs, "\n")
		for i := 1; i <= 3; i++ {
			if err := cli.Run([]string{"migrator", "up"}); err != nil {
				t.Fatal(err)
			}
		}
		if err := cli.Run([]string{"migrator", "latest"}); err != nil {
			t.Fatal(err)
		}
	})
}

func TestMultiplesQueriesInSameMigraition(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/e2e/mocks/2")
	env(func(envs env_set) {
		t.Log("\nTesting with envs: ", envs, "\n")
		if err := cli.Run([]string{"migrator", "latest"}); err != nil {
			t.Fatal(err)
		}
	})
}
