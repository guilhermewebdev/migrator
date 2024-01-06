package e2e_tests

import (
	"testing"

	"github.com/guilhermewebdev/migrator/src/cli"
)

func TestUnlock(t *testing.T) {
	env(func(envs env_set) {
		t.Log("Testing with envs: ", envs)
		if err := cli.Run([]string{"migrator", "unlock"}); err != nil {
			t.Fatal(err)
		}
	})
}
