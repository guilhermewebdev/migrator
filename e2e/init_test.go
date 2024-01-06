package e2e_tests

import (
	"os"
	"testing"

	"github.com/guilhermewebdev/migrator/src/cli"
)

func TestInit(t *testing.T) {
	conf_file := os.Getenv("ROOT_DIR") + "/tmp/tmp-migrator.yml"
	defer os.Remove(conf_file)
	if err := cli.Run([]string{"migrator", "-c", conf_file, "init"}); err != nil {
		t.Fatal(err)
	}
	if !file_exists(conf_file) {
		t.Fatalf("%s file was not created", conf_file)
	}
}
