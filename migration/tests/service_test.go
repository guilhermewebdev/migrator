package migration_test

import (
	"regexp"
	"testing"

	"github.com/guilhermewebdev/migrator/migration"
)

type referenceRepositoryMock struct {
	listMock              []migration.Reference
	listMockError         error
	migrationsRan         []migration.Migration
	migrationRunMockError error
	lockStatus            bool
	lockMockError         error
	prepareMockError      error
}

func (r *referenceRepositoryMock) Prepare() error {
	return r.prepareMockError
}

func (r *referenceRepositoryMock) List() ([]migration.Reference, error) {
	return r.listMock, r.listMockError
}

func (r *referenceRepositoryMock) Run(migration migration.Migration) error {
	r.migrationsRan = append(r.migrationsRan, migration)
	return r.migrationRunMockError
}

func (r *referenceRepositoryMock) Lock() error {
	r.lockStatus = true
	return r.lockMockError
}

func (r *referenceRepositoryMock) Unlock() error {
	r.lockStatus = false
	return r.lockMockError
}

func (r *referenceRepositoryMock) IsLocked() (bool, error) {
	return r.lockStatus, r.lockMockError
}

type migrationsRepositoryMock struct {
	creations         []string
	creationMockError error
	listMock          []migration.Migration
	listMockError     error
}

func (r *migrationsRepositoryMock) Create(name string) error {
	r.creations = append(r.creations, name)
	return r.creationMockError
}

func (r *migrationsRepositoryMock) List() ([]migration.Migration, error) {
	return r.listMock, r.listMockError
}

func TestCreate(t *testing.T) {
	migrations := &migrationsRepositoryMock{}
	references := &referenceRepositoryMock{}
	var service migration.Service = &migration.ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	err := service.Create("create_user")
	if err != nil {
		t.Error(err)
	}
	pattern := `^[0-9]{1,}_[A-z_]{1,}$`
	matched0, err := regexp.MatchString(pattern, migrations.creations[0])
	if err != nil || !matched0 {
		t.Log(migrations.creations, matched0)
		t.Fail()
	}
}
