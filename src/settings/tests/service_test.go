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

func (s *settingsRepositoryMock) GetFromEnv(initial stgs.Settings) (stgs.Settings, error) {
	return initial, s.getFromEnvError
}

func (s *settingsRepositoryMock) GetFromFile(_ string, initial stgs.Settings) (stgs.Settings, error) {
	return initial, s.getFromFileError
}

func TestGetSettings_Default(t *testing.T) {
	var repository stgs.SettingsRepository = &settingsRepositoryMock{
		getFromEnvResponse:  stgs.Settings{},
		getFromFileResponse: stgs.Settings{},
		getFromEnvError:     nil,
		getFromFileError:    nil,
	}
	service := &stgs.SettingsServiceImpl{
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
