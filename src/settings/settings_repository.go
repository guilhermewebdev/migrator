package settings

import (
	"os"

	"github.com/guilhermewebdev/migrator/src/lib"
	"gopkg.in/yaml.v2"
)

type SettingsRepository interface {
	GetFromEnv() (Settings, error)
	GetFromFile(file_name string) (Settings, error)
	CreateFile(file_name string) error
	WriteFile(file_name string, content string) error
}

type SettingsRepositoryImpl struct {
	Disk lib.Disk
}

func (r *SettingsRepositoryImpl) GetFromEnv() (Settings, error) {
	var settings Settings = Settings{
		DB_DSN:              os.Getenv("DB_DSN"),
		DB_Driver:           os.Getenv("DB_DRIVER"),
		MigrationsDir:       os.Getenv("MIGRATIONS_DIR"),
		MigrationsTableName: os.Getenv("MIGRATIONS_TABLE"),
	}
	return settings, nil
}

func (r *SettingsRepositoryImpl) get_settings_file_content(file_path string) (Settings, error) {
	data, err := r.Disk.Read(file_path)
	settings := Settings{}
	if err != nil {
		return settings, err
	}
	data_with_envs := os.ExpandEnv(data)
	err = yaml.Unmarshal([]byte(data_with_envs), &settings)
	if err != nil {
		return settings, err
	}
	return settings, nil
}

func (r *SettingsRepositoryImpl) GetFromFile(file_name string) (Settings, error) {
	file_path, err := r.Disk.SearchFileInParentDirectories(file_name)
	empty := Settings{}
	if err != nil || file_path == "" {
		return empty, err
	}
	return r.get_settings_file_content(file_path)
}

func (r *SettingsRepositoryImpl) CreateFile(file_name string) error {
	return nil
}

func (r *SettingsRepositoryImpl) WriteFile(file_name string, content string) error {
	return nil
}
