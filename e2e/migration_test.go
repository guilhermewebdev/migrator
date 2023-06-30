package e2e_tests

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/guilhermewebdev/migrator/cli"
)

type env_set map[string]string

func set_env(envs env_set) {
	for k, v := range envs {
		os.Setenv(k, v)
	}
}

var test_envs = []env_set{
	{
		"DB_DSN":    "user:pass@tcp(mysql:3306)/test",
		"DB_DRIVER": "mysql",
	},
	{
		"DB_DSN":    "postgres://user:pass@postgres:5432/test?sslmode=disable",
		"DB_DRIVER": "postgres",
	},
	{
		"DB_DSN":    "/usr/src/migrator/tmp/test.sqlite3",
		"DB_DRIVER": "sqlite3",
	},
}

func env(test func(env_set)) {
	table_name := "table_" + strings.Replace(uuid.NewString(), "-", "_", 4)
	os.Setenv("MIGRATIONS_TABLE", table_name)
	file, err := os.OpenFile("/usr/src/migrator/tmp/test.sqlite3", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	count := 1
	for _, envs := range test_envs {
		set_env(envs)
		test(envs)
		count++
	}
}

func TestUp(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", "/usr/src/migrator/e2e/mocks/1")
	env(func(envs env_set) {
		t.Log("Testing with envs: ", envs)
		if err := cli.Run([]string{"migrator", "up"}); err != nil {
			t.Error(err)
		}
	})
}

func TestNew(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", "/usr/src/migrator/tmp/migrations")
	if err := cli.Run([]string{"migrator", "new", "migration_test"}); err != nil {
		t.Error(err)
	}
}

func TestUnlock(t *testing.T) {
	env(func(envs env_set) {
		t.Log("Testing with envs: ", envs)
		if err := cli.Run([]string{"migrator", "unlock"}); err != nil {
			t.Error(err)
		}
	})
}
