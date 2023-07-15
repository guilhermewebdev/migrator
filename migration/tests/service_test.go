package migration_test

import (
	"fmt"
	"regexp"
	"testing"

	mod "github.com/guilhermewebdev/migrator/migration"
)

type referenceRepositoryMock struct {
	listMock              []mod.Reference
	listMockError         error
	migrationsRan         []mod.Migration
	migrationRunMockError error
	referenceMock         mod.Reference
	referenceMockError    error
	lockStatus            bool
	lockMockError         error
	prepareMockError      error
}

func (r *referenceRepositoryMock) Prepare() error {
	return r.prepareMockError
}

func (r *referenceRepositoryMock) List() ([]mod.Reference, error) {
	return r.listMock, r.listMockError
}

func (r *referenceRepositoryMock) Up(migration mod.Migration) error {
	r.migrationsRan = append(r.migrationsRan, migration)
	return r.migrationRunMockError
}

func (r *referenceRepositoryMock) Down(migration mod.Migration) error {
	r.migrationsRan = append(r.migrationsRan, migration)
	return r.migrationRunMockError
}

func (r *referenceRepositoryMock) GetLast() (mod.Reference, error) {
	return r.referenceMock, r.referenceMockError
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
	creations          []string
	creationMockError  error
	listMock           []mod.Migration
	listMockError      error
	migrationMock      mod.Migration
	migrationErrorMock error
}

func (r *migrationsRepositoryMock) Create(name string) error {
	r.creations = append(r.creations, name)
	return r.creationMockError
}

func (r *migrationsRepositoryMock) List() ([]mod.Migration, error) {
	return r.listMock, r.listMockError
}

func (r *migrationsRepositoryMock) Read(name string) (mod.Migration, error) {
	return r.migrationMock, r.migrationErrorMock
}

func TestService_Create(t *testing.T) {
	t.Parallel()
	migrations := &migrationsRepositoryMock{}
	references := &referenceRepositoryMock{}
	var service mod.Service = &mod.ServiceImpl{
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
		t.Fatal()
	}
}

func TestService_Up_WithPendingMigration(t *testing.T) {
	t.Parallel()
	pending_migration := mod.Migration{
		Name: "testing",
		Path: "testing",
	}
	migrations := &migrationsRepositoryMock{
		listMock: []mod.Migration{pending_migration},
	}
	references := &referenceRepositoryMock{}
	var service mod.Service = &mod.ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	migration, err := service.Up()
	if err != nil {
		t.Error(err)
	}
	if migration != pending_migration {
		t.Fatal()
	}
}

func TestService_Up_WhenAllMigrationsAreRan(t *testing.T) {
	t.Parallel()
	migrations := &migrationsRepositoryMock{
		listMock: []mod.Migration{{
			Name: "testing",
			Path: "testing",
		}},
	}
	references := &referenceRepositoryMock{
		listMock: []mod.Reference{{
			Name: "testing",
		}},
	}
	var service mod.Service = &mod.ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	ran, err := service.Up()
	if err != nil {
		t.Error(err)
	}
	expected := mod.Migration{}
	if ran != expected {
		t.Fatal()
	}
}

func TestService_Up_WhenHasNoMigrations(t *testing.T) {
	t.Parallel()
	migrations := &migrationsRepositoryMock{}
	references := &referenceRepositoryMock{}
	var service mod.Service = &mod.ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	ran, err := service.Up()
	if err != nil {
		t.Error(err)
	}
	expected := mod.Migration{}
	if ran != expected {
		t.Fatal()
	}
}

func TestService_Up_WhenMigrationsAreCorrupted(t *testing.T) {
	t.Parallel()
	expected_migration := mod.Migration{}
	migrations := &migrationsRepositoryMock{
		listMock: []mod.Migration{
			{
				Name: "testing",
				Path: "testing",
			},
		},
	}
	references := &referenceRepositoryMock{
		listMock: []mod.Reference{{
			Name: "wrong_testing",
		}},
	}
	var service mod.Service = &mod.ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	ran, err := service.Up()
	if err == nil {
		t.Fatal("No one error are raised")
	}
	if ran != expected_migration {
		t.Fatal(ran, " is not ", expected_migration)
	}
	if err.Error() != "Migrations are corrupted" {
		t.Fatal(err)
	}
}

func TestService_Up_WhenMigrationsAreDisorderly(t *testing.T) {
	t.Parallel()
	migrations := &migrationsRepositoryMock{
		listMock: []mod.Migration{
			{
				Name: "0_testing",
				Path: "testing",
			},
			{
				Name: "1_testing",
				Path: "testing",
			},
			{
				Name: "2_testing",
				Path: "testing",
			},
		},
	}
	references := &referenceRepositoryMock{
		listMock: []mod.Reference{
			{
				Name: "0_testing",
			},
			{
				Name: "2_testing",
			},
		},
	}
	var service mod.Service = &mod.ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	ran, err := service.Up()
	if err != nil {
		t.Fatal(err)
	}
	expected := mod.Migration{
		Name: "1_testing",
		Path: "testing",
	}
	if ran != expected {
		t.Fatal(ran, " is not ", expected)
	}
}

func TestService_Unlock(t *testing.T) {
	t.Parallel()
	migrations := &migrationsRepositoryMock{}
	references := &referenceRepositoryMock{}
	var service mod.Service = &mod.ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	if err := service.Unlock(); err != nil {
		t.Fatal(err)
	}
}

func TestService_Down(t *testing.T) {
	t.Parallel()
	expected_migration := mod.Migration{Name: "2_testing", Path: "testing"}
	migrations := &migrationsRepositoryMock{
		migrationMock: expected_migration,
	}
	references := &referenceRepositoryMock{
		referenceMock: mod.Reference{
			Name: "2_testing",
		},
	}
	var service mod.Service = &mod.ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	migration, err := service.Down()
	if err != nil {
		t.Fatal(err)
	}
	if migration != expected_migration {
		t.Fatal(migration, " is not ", expected_migration)
	}
}

func TestService_Down_WithoutReferences(t *testing.T) {
	t.Parallel()
	migrations := &migrationsRepositoryMock{
		migrationMock: mod.Migration{Name: "2_testing", Path: "testing"},
	}
	references := &referenceRepositoryMock{
		referenceMockError: fmt.Errorf("No migrations to rollback"),
	}
	var service mod.Service = &mod.ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	migration, err := service.Down()
	if err == nil || err.Error() != "No migrations to rollback" {
		t.Fatal(err)
	}
	empty := mod.Migration{}
	if migration != empty {
		t.Fatal(migration, " is not ", empty)
	}
}
