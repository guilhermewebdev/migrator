package settings_test

import (
	"testing"

	stgs "github.com/guilhermewebdev/migrator/src/settings"
)

type settingsRepositoryMock struct {
	getFromEnvResponse  stgs.Settings
	getFromFileResponse stgs.Settings
	getFromEnvError     error
	getFromFileError    error
}

func (s *settingsRepositoryMock) GetFromEnv() (stgs.Settings, error) {
	return s.getFromEnvResponse, s.getFromEnvError
}

func (s *settingsRepositoryMock) GetFromFile(_ string) (stgs.Settings, error) {
	return s.getFromFileResponse, s.getFromFileError
}

func TestGetSettings_Default(t *testing.T) {
	var repository stgs.SettingsRepository = &settingsRepositoryMock{
		getFromEnvResponse:  stgs.Settings{},
		getFromFileResponse: stgs.Settings{},
		getFromEnvError:     nil,
		getFromFileError:    nil,
	}
	service := &stgs.ServiceImpl{
		Settings: repository,
	}
	settings, err := service.Get("migrator.yml")
	if err != nil {
		t.Fatal(err)
	}
	expected := stgs.Settings{
		MigrationsDir:       "./migrations",
		MigrationsTableName: "migrations",
		DB_DSN:              "",
		DB_Driver:           "",
	}
	if settings != expected {
		t.Fatal(expected, "is not", settings)
	}
}

func TestGetSettings_WhenMigrationsDirComesFromFile(t *testing.T) {
	var repository stgs.SettingsRepository = &settingsRepositoryMock{
		getFromEnvResponse: stgs.Settings{},
		getFromFileResponse: stgs.Settings{
			MigrationsDir: "./m",
		},
		getFromEnvError:  nil,
		getFromFileError: nil,
	}
	service := &stgs.ServiceImpl{
		Settings: repository,
	}
	settings, err := service.Get("migrator.yml")
	if err != nil {
		t.Fatal(err)
	}
	expected := stgs.Settings{
		MigrationsDir:       "./m",
		MigrationsTableName: "migrations",
		DB_DSN:              "",
		DB_Driver:           "",
	}
	if settings != expected {
		t.Fatal(expected, "is not", settings)
	}
}

func TestGetSettings_WhenMigrationsDirComesFromEnv(t *testing.T) {
	var repository stgs.SettingsRepository = &settingsRepositoryMock{
		getFromEnvResponse: stgs.Settings{
			MigrationsDir: "./m",
		},
		getFromFileResponse: stgs.Settings{},
		getFromEnvError:     nil,
		getFromFileError:    nil,
	}
	service := &stgs.ServiceImpl{
		Settings: repository,
	}
	settings, err := service.Get("migrator.yml")
	if err != nil {
		t.Fatal(err)
	}
	expected := stgs.Settings{
		MigrationsDir:       "./m",
		MigrationsTableName: "migrations",
		DB_DSN:              "",
		DB_Driver:           "",
	}
	if settings != expected {
		t.Fatal(expected, "is not", settings)
	}
}

func TestGetSettings_WhenDB_DSNComesFromEnv(t *testing.T) {
	var repository stgs.SettingsRepository = &settingsRepositoryMock{
		getFromEnvResponse: stgs.Settings{
			MigrationsDir: "./m",
			DB_DSN:        "postgres://host..",
		},
		getFromFileResponse: stgs.Settings{},
		getFromEnvError:     nil,
		getFromFileError:    nil,
	}
	service := &stgs.ServiceImpl{
		Settings: repository,
	}
	settings, err := service.Get("migrator.yml")
	if err != nil {
		t.Fatal(err)
	}
	expected := stgs.Settings{
		MigrationsDir:       "./m",
		MigrationsTableName: "migrations",
		DB_DSN:              "postgres://host..",
		DB_Driver:           "",
	}
	if settings != expected {
		t.Fatal(expected, "is not", settings)
	}
}
