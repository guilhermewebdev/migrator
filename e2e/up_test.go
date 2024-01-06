package e2e_tests

import (
	"os"
	"testing"

	"github.com/guilhermewebdev/migrator/src/cli"
)

func TestUp(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/e2e/mocks/1")
	env(func(envs env_set) {
		t.Log("Testing with envs: ", envs)
		if err := cli.Run([]string{"migrator", "up"}); err != nil {
			t.Fatal(err)
		}
	})
}
