package e2e_tests

import (
	"os"
	"testing"

	"github.com/guilhermewebdev/migrator/src/cli"
)

func TestNew(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/tmp/migrations")
	if err := cli.Run([]string{"migrator", "new", "migration_test"}); err != nil {
		t.Fatal(err)
	}
}
