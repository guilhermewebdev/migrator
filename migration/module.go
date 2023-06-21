package migration

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type MigrationModule interface {
	Controller() Controller
}

type MigrationModuleImpl struct {
	controller Controller
}

func (mod *MigrationModuleImpl) Controller() Controller {
	return mod.controller
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
	migrations_dir := "./migrations"
	current_dir, err := filepath.Abs(migrations_dir)
	if err != nil {
		current_dir = migrations_dir
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
		controller: controller,
	}
	return module
}
