package e2e_tests

import (
	"os"
	"strings"
	"testing"

	"github.com/guilhermewebdev/migrator/src/cli"
)

func TestUp(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/e2e/mocks/1")
	env(func(envs env_set) {
		output, err := capture_output(func() error {
			t.Log("Testing with envs: ", envs)
			return cli.Run([]string{"migrator", "up"})
		})
		if err != nil {
			t.Fatal(err)
		}
		expected_log := "The migration \"1688099609991_test_init\" was ran."
		output = strings.Trim(output, "\n")
		if output != expected_log {
			t.Error(output, "is not", expected_log)
			t.Fatal("Invalid migration order execution")
		}
	})
}
