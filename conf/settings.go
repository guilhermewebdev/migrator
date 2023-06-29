package conf

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	MigrationsDir       string `yaml:"migrations_dir"`
	MigrationsTableName string `yaml:"migrations_table_name"`
	DBDSN               string `yaml:"db_dsn"`
	DBDriver            string `yaml:"db_driver"`
}

func search_file_in_parent_directories(file_name string) (string, error) {
	current_dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		file_path := filepath.Join(current_dir, file_name)
		_, err := os.Stat(file_path)
		if err == nil {
			return file_path, nil
		}
		if current_dir == filepath.Dir(current_dir) {
			break
		}
		current_dir = filepath.Dir(current_dir)
	}
	return "", fmt.Errorf("Arquivo '%s' não encontrado nos diretórios pais", file_name)
}

func get_initial_settings() Settings {
	var settings Settings = Settings{
		DBDSN:         os.Getenv("DB_DSN"),
		DBDriver:      os.Getenv("DB_DRIVER"),
		MigrationsDir: os.Getenv("MIGRATIONS_DIR"),
	}
	return settings
}

func LoadSettings(settings_file_name string) Settings {
	initial_settings := get_initial_settings()
	settings_path, err := search_file_in_parent_directories(settings_file_name)
	if err != nil {
		return get_default_settings(initial_settings)
	}
	settings, err := get_settings_file_content(settings_path, initial_settings)
	if err != nil {
		log.Fatal(err)
	}
	return settings
}

func get_default_settings(settings Settings) Settings {
	migrations_dir := "./migrations"
	current_dir, err := filepath.Abs(migrations_dir)
	if err != nil {
		current_dir = migrations_dir
	}
	settings.MigrationsDir = current_dir
	return settings
}

func get_settings_file_content(file_path string, settings Settings) (Settings, error) {
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		return settings, err
	}
	data_with_envs := os.ExpandEnv(string(data))
	err = yaml.Unmarshal([]byte(data_with_envs), &settings)
	if err != nil {
		return settings, err
	}
	return settings, nil
}
