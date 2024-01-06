package settings_test

import (
	"os"
	"testing"

	lib_mocks "github.com/guilhermewebdev/migrator/src/lib/mocks"
	"github.com/guilhermewebdev/migrator/src/settings"
)

func TestGetSettingsFromEnv(t *testing.T) {
	var repository settings.SettingsRepository = &settings.SettingsRepositoryImpl{
		Disk: &lib_mocks.DiskMock{},
	}
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
	var repository settings.SettingsRepository = &settings.SettingsRepositoryImpl{
		Disk: &lib_mocks.DiskMock{},
	}
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
	var repository settings.SettingsRepository = &settings.SettingsRepositoryImpl{
		Disk: &lib_mocks.DiskMock{},
	}
	expected := settings.Settings{}
	settings, err := repository.GetFromFile("./mocks/migrator-wrong.yml")
	if err == nil {
		t.Fatal("Error should be raised")
	}
	if settings != expected {
		t.Fatal(settings, "is not", expected)
	}
}

func TestGetSettingsFromFile_WhenFileNotExists(t *testing.T) {
	var repository settings.SettingsRepository = &settings.SettingsRepositoryImpl{
		Disk: &lib_mocks.DiskMock{},
	}
	expected := settings.Settings{}
	settings, err := repository.GetFromFile("migrator.yml")
	if err != nil {
		t.Fatal("Error should be raised")
	}
	if settings != expected {
		t.Fatal(settings, "is not", expected)
	}
}

func TestCreateFile(t *testing.T) {
	disk := &lib_mocks.DiskMock{}
	var repository settings.SettingsRepository = &settings.SettingsRepositoryImpl{
		Disk: disk,
	}
	err := repository.CreateFile("migrator.yml")
	if err != nil {
		t.Fatal(err)
	}
	if disk.Creations[0] != "migrator.yml" {
		t.Fatal("File was not created")
	}
	if len(disk.Creations) != 1 {
		t.Fatal("Different than 1 files were created")
	}
}

func TestWriteFile(t *testing.T) {
	disk := &lib_mocks.DiskMock{}
	var repository settings.SettingsRepository = &settings.SettingsRepositoryImpl{
		Disk: disk,
	}
	err := repository.WriteFile("migrator.yml", "testing content")
	if err != nil {
		t.Fatal(err)
	}
	if disk.Writes[0][0] != "migrator.yml" {
		t.Fatal("File was not created")
	}
	if disk.Writes[0][1] != "testing content" {
		t.Fatal("File was not wrote")
	}
}
