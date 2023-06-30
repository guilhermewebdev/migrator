package migration

import (
	"log"
	"path"
	"regexp"
	"testing"

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

func TestMigrationRepository_Create(t *testing.T) {
	disk := diskMock{}
	var repo MigrationRepository = &MigrationRepositoryImpl{
		Disk:     &disk,
		Settings: get_settings(),
	}
	if err := repo.Create("test"); err != nil {
		t.Fatal(err)
	}
	if matched, err := regexp.Match(
		"test/(up|down).sql$",
		[]byte(disk.Creations[0]),
	); err != nil || !matched {
		log.Fatal(err, matched, disk.Creations)
	}
}

func TestMigrationRepository_List(t *testing.T) {
	disk := diskMock{
		listMock: []string{
			"test",
		},
		readMock: "SELECT * FROM table;",
	}
	var repo MigrationRepository = &MigrationRepositoryImpl{
		Disk:     &disk,
		Settings: get_settings(),
	}
	migrations, err := repo.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(migrations) != len(disk.listMock) {
		t.Fatal(migrations, disk)
	}
	if migrations[0].UpQuery != disk.readMock {
		t.Fatal(migrations, disk)
	}
}
