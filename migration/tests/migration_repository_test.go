package migration_test

import (
	"path"

	"github.com/guilhermewebdev/migrator/conf"
)

type diskMock struct {
	Creations     []string
	Lists         []string
	Reads         []string
	creationMock  error
	listMock      []string
	listErrorMock error
	readMock      string
	readErrorMock error
}

func (repo *diskMock) Create(path_name string, file_name string) error {
	full_path := path.Join(path_name, file_name)
	repo.Creations = append(repo.Creations, full_path)
	return repo.creationMock
}

func (repo *diskMock) List(dir string) ([]string, error) {
	repo.Lists = append(repo.Lists, dir)
	return repo.listMock, repo.listErrorMock
}

func (repo *diskMock) Read(file_path string) (string, error) {
	repo.Reads = append(repo.Reads, file_path)
	return repo.readMock, repo.readErrorMock
}

func get_settings() conf.Settings {
	settings := conf.Settings{
		MigrationsDir: "./migrations",
	}
	return settings
}
