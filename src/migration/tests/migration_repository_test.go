package migration_test

import (
	"log"
	"regexp"
	"testing"

	lib_mocks "github.com/guilhermewebdev/migrator/src/lib/mocks"
	"github.com/guilhermewebdev/migrator/src/migration"
	"github.com/guilhermewebdev/migrator/src/settings"
)

func get_settings() settings.Settings {
	settings := settings.Settings{
		MigrationsDir:       "./migrations",
		MigrationsTableName: "migrations",
	}
	return settings
}

func TestMigrationRepository_Create(t *testing.T) {
	t.Parallel()
	disk := lib_mocks.DiskMock{}
	var repo migration.MigrationRepository = &migration.MigrationRepositoryImpl{
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
	t.Parallel()
	disk := lib_mocks.DiskMock{
		ListMock: []string{
			"test",
		},
		ReadMock: "SELECT * FROM table;",
	}
	var repo migration.MigrationRepository = &migration.MigrationRepositoryImpl{
		Disk:     &disk,
		Settings: get_settings(),
	}
	migrations, err := repo.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(migrations) != len(disk.ListMock) {
		t.Fatal(migrations, disk)
	}
	if migrations[0].UpQuery != disk.ReadMock {
		t.Fatal(migrations, disk)
	}
}
