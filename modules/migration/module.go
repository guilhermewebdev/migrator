package migration

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type MigrationModule interface{}

type MigrationModuleImpl struct {
	Controller Controller
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

func get_default_settings() Settings {
	current_dir, err := filepath.Abs(".")
	if err != nil {
		current_dir = "."
	}
	return Settings{
		MigrationsDir: current_dir,
	}
}

func get_settings_file_content(file_path string) Settings {
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(err)
	}
	var settings Settings
	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		log.Fatal(err)
	}
	return settings
}

func load_settings(settings_file_name string) Settings {
	settings_path, err := search_file_in_parent_directories(settings_file_name)
	if err != nil {
		return get_default_settings()
	}
	settings := get_settings_file_content(settings_path)
	return settings
}

func NewMigrationModule() MigrationModule {
	var disk DiskRepository = &DiskRepositoryImpl{}
	var service Service = &ServiceImpl{
		Disk:     disk,
		Settings: load_settings("migrator.yml"),
	}
	var controller Controller = &ControllerImpl{
		Service: service,
	}
	var module MigrationModule = &MigrationModuleImpl{
		Controller: controller,
	}
	return module
}
