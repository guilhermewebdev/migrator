package settings_test

import (
	"os"
	"testing"

	"github.com/guilhermewebdev/migrator/src/settings"
)

func TestGetSettingsFromEnv(t *testing.T) {
	var repository settings.SettingsRepository = &settings.SettingsRepositoryImpl{}
	os.Setenv("DB_DSN", "a")
	os.Setenv("DB_DRIVER", "b")
	os.Setenv("MIGRATIONS_DIR", "c")
	os.Setenv("MIGRATIONS_TABLE", "d")
	expected := settings.Settings{
		MigrationsDir:       "c",
		MigrationsTableName: "d",
		DB_DSN:              "a",
		DB_Driver:           "b",
	}
	settings, err := repository.GetFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if settings != expected {
		t.Fatal(settings, "is not", expected)
	}
}

func TestGetSettingsFromFile(t *testing.T) {
	var repository settings.SettingsRepository = &settings.SettingsRepositoryImpl{}
	os.Setenv("DB_DSN", "a")
	os.Setenv("DB_DRIVER", "b")
	os.Setenv("MIGRATIONS_DIR", "c")
	os.Setenv("MIGRATIONS_TABLE", "d")
	expected := settings.Settings{
		MigrationsDir:       "./test",
		MigrationsTableName: "testing_table_name",
		DB_DSN:              "testing_dsn",
		DB_Driver:           "testing_driver",
	}
	settings, err := repository.GetFromFile("./mocks/migrator.yml")
	if err != nil {
		t.Fatal(err)
	}
	if settings != expected {
		t.Fatal(settings, "is not", expected)
	}
}

func TestGetSettingsFromFile_WhenFileIsInvalid(t *testing.T) {
	var repository settings.SettingsRepository = &settings.SettingsRepositoryImpl{}
	expected := settings.Settings{}
	settings, err := repository.GetFromFile("./mocks/migrator-wrong.yml")
	if err == nil {
		t.Fatal("Error should be raised")
	}
	if settings != expected {
		t.Fatal(settings, "is not", expected)
	}
}
