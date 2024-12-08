package e2e_tests

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
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
		"DB_DSN":    "user:pass@tcp(mysql:3306)/test?multiStatements=true",
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

func capture_output(f func() error) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	err := f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out), err
}
