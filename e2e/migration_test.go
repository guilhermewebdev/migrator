package e2e_tests

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/guilhermewebdev/migrator/src/cli"
)

type env_set map[string]string

func set_env(envs env_set) {
	for k, v := range envs {
		os.Setenv(k, v)
	}
}

func file_exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
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
		"DB_DSN":    os.Getenv("ROOT_DIR") + "/tmp/test.sqlite3",
		"DB_DRIVER": "sqlite3",
	},
	{
		"DB_DSN":    os.Getenv("ROOT_DIR") + "/tmp/test.sqlite",
		"DB_DRIVER": "sqlite",
	},
}

func new_sqlite(file_path string) {
	file, err := os.OpenFile(file_path, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func env(test func(env_set)) {
	table_name := "table_" + strings.Replace(uuid.NewString(), "-", "_", 4)
	os.Setenv("MIGRATIONS_TABLE", table_name)
	new_sqlite(os.Getenv("ROOT_DIR") + "/tmp/test.sqlite")
	new_sqlite(os.Getenv("ROOT_DIR") + "/tmp/test.sqlite3")
	count := 1
	for _, envs := range test_envs {
		set_env(envs)
		test(envs)
		count++
	}
}

func TestUp(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/e2e/mocks/1")
	env(func(envs env_set) {
		t.Log("Testing with envs: ", envs)
		if err := cli.Run([]string{"migrator", "up"}); err != nil {
			t.Fatal(err)
		}
	})
}

func TestNew(t *testing.T) {
	os.Setenv("MIGRATIONS_DIR", os.Getenv("ROOT_DIR")+"/tmp/migrations")
	if err := cli.Run([]string{"migrator", "new", "migration_test"}); err != nil {
		t.Fatal(err)
	}
}

func TestUnlock(t *testing.T) {
	env(func(envs env_set) {
		t.Log("Testing with envs: ", envs)
		if err := cli.Run([]string{"migrator", "unlock"}); err != nil {
			t.Fatal(err)
		}
	})
}

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
