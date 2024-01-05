package settings

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type SettingsRepository interface {
	GetFromEnv() (Settings, error)
	GetFromFile(file_name string) (Settings, error)
}

type SettingsRepositoryImpl struct{}

func (r *SettingsRepositoryImpl) GetFromEnv() (Settings, error) {
	var settings Settings = Settings{
		DB_DSN:              os.Getenv("DB_DSN"),
		DB_Driver:           os.Getenv("DB_DRIVER"),
		MigrationsDir:       os.Getenv("MIGRATIONS_DIR"),
		MigrationsTableName: os.Getenv("MIGRATIONS_TABLE"),
	}
	return settings, nil
}

func (r *SettingsRepositoryImpl) search_file_in_parent_directories(file_name string) (string, error) {
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
	return "", fmt.Errorf("File '%s' not found on parent directories", file_name)
}

func (r *SettingsRepositoryImpl) get_settings_file_content(file_path string) (Settings, error) {
	data, err := os.ReadFile(file_path)
	settings := Settings{}
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

func (r *SettingsRepositoryImpl) GetFromFile(file_name string) (Settings, error) {
	file_path, err := r.search_file_in_parent_directories(file_name)
	empty := Settings{}
	if err != nil {
		return empty, err
	}
	return r.get_settings_file_content(file_path)
}
