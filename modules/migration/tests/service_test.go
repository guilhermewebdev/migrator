package migration_test

import (
	"path"
	"regexp"
	"testing"

	"github.com/guilhermewebdev/migrator/modules/migration"
)

type diskRepositoryMock struct {
	Creations     []string
	Lists         []string
	Reads         []string
	creationMock  error
	listMock      []string
	listErrorMock error
	readMock      string
	readErrorMock error
}

func (repo *diskRepositoryMock) Create(path_name string, file_name string) error {
	full_path := path.Join(path_name, file_name)
	repo.Creations = append(repo.Creations, full_path)
	return repo.creationMock
}

func (repo *diskRepositoryMock) List(dir string) ([]string, error) {
	repo.Lists = append(repo.Lists, dir)
	return repo.listMock, repo.listErrorMock
}

func (repo *diskRepositoryMock) Read(file_path string) (string, error) {
	repo.Reads = append(repo.Reads, file_path)
	return repo.readMock, repo.readErrorMock
}

func get_settings() migration.Settings {
	settings := migration.Settings{
		MigrationsDir: "./migrations",
	}
	return settings
}

func get_service(diskMock migration.DiskRepository) migration.Service {
	var service migration.Service = &migration.ServiceImpl{
		Disk:     diskMock,
		Settings: get_settings(),
	}
	return service
}

func TestCreate(t *testing.T) {
	var disk = &diskRepositoryMock{}
	service := get_service(disk)
	err := service.Create("create_user")
	if err != nil {
		t.Error(err)
	}
	pattern := `^migrations/[0-9]{1,}_[A-z_]{1,}/(up|down).sql$`
	matched0, err := regexp.MatchString(pattern, disk.Creations[0])
	if err != nil || !matched0 {
		t.Log(disk.Creations, matched0)
		t.Fail()
	}
	matched1, err := regexp.MatchString(pattern, disk.Creations[1])
	if err != nil || !matched1 {
		t.Log(disk.Creations, matched1)
		t.Fail()
	}
}
