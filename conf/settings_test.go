package conf

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/guilhermewebdev/migrator/lib"
)

func mock_file(path_name string, file_name string) func() {
	disk := lib.DiskImpl{}
	disk.Create(path_name, file_name)
	return func() {
		os.Remove(path.Join(path_name, file_name))
	}
}

func TestSearchFileInParentDirectories(t *testing.T) {
	defer mock_file("..", "conf-temp.yml")()
	file_path, err := search_file_in_parent_directories("conf-temp.yml")
	if err != nil {
		t.Error(err)
	}
	abs_path, err := filepath.Abs(path.Join("..", "conf-temp.yml"))
	if file_path != abs_path {
		t.Fail()
	}
}

func TestSearchFileInParentDirectories_AndNotFindIt(t *testing.T) {
	_, err := search_file_in_parent_directories("conf-temp.yml")
	if err == nil {
		t.Fail()
	}
	if err.Error() != "Arquivo 'conf-temp.yml' não encontrado nos diretórios pais" {
		t.Fail()
	}
}

func TestGetDefaultSettings(t *testing.T) {
	migrations_dir, _ := filepath.Abs("./migrations")
	expected_settings := Settings{
		MigrationsDir: migrations_dir,
	}
	received_settings := get_default_settings(get_initial_settings())
	if received_settings != expected_settings {
		t.Log(received_settings, " is not ", expected_settings)
		t.Fail()
	}
}

func TestGetSettingsFileContent(t *testing.T) {
	settings, _ := get_settings_file_content("./mocks/migrator.yml", get_initial_settings())
	expected_settings := Settings{
		MigrationsDir:       "./test",
		MigrationsTableName: "testing_table_name",
		DBDSN:               "testing_dsn",
		DBDriver:            "testing_driver",
	}
	if settings != expected_settings {
		t.Log(settings, " is not ", expected_settings)
		t.Fail()
	}
}

func TestGetSettingsFileContent_WithWrongSintaxe(t *testing.T) {
	_, err := get_settings_file_content("./mocks/migrator-wrong.yml", get_initial_settings())
	if err == nil {
		t.Fail()
	}
}

func TestLoadSettings_WithFile(t *testing.T) {
	settings_from_file, _ := get_settings_file_content("./mocks/migrator.yml", get_initial_settings())
	settings := LoadSettings("./mocks/migrator.yml")
	if settings_from_file != settings {
		t.Log(settings, " is not ", settings_from_file)
	}
}

func TestLoadSettings_WithDefault(t *testing.T) {
	default_settings := get_default_settings(get_initial_settings())
	settings := LoadSettings("unexistent.file")
	if default_settings != settings {
		t.Log(settings, " is not ", default_settings)
	}
}

func TestGetInitialSettings(t *testing.T) {
	os.Setenv("DB_DSN", "testing_dsn")
	os.Setenv("DB_DRIVER", "testing_driver")
	settings := get_initial_settings()
	expected_settings := Settings{
		DBDSN:    "testing_dsn",
		DBDriver: "testing_driver",
	}
	if settings != expected_settings {
		t.Log(settings, " is not ", expected_settings)
		t.Fail()
	}
}
