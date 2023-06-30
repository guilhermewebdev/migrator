package migration

import (
	"regexp"
	"testing"
)

type referenceRepositoryMock struct {
	listMock              []Reference
	listMockError         error
	migrationsRan         []Migration
	migrationRunMockError error
	lockStatus            bool
	lockMockError         error
	prepareMockError      error
}

func (r *referenceRepositoryMock) Prepare() error {
	return r.prepareMockError
}

func (r *referenceRepositoryMock) List() ([]Reference, error) {
	return r.listMock, r.listMockError
}

func (r *referenceRepositoryMock) Run(migration Migration) error {
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
	listMock          []Migration
	listMockError     error
}

func (r *migrationsRepositoryMock) Create(name string) error {
	r.creations = append(r.creations, name)
	return r.creationMockError
}

func (r *migrationsRepositoryMock) List() ([]Migration, error) {
	return r.listMock, r.listMockError
}

func TestCreate(t *testing.T) {
	migrations := &migrationsRepositoryMock{}
	references := &referenceRepositoryMock{}
	var service Service = &ServiceImpl{
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

func TestUp_WithPendingMigration(t *testing.T) {
	pending_migration := Migration{
		Name: "testing",
		Path: "testing",
	}
	migrations := &migrationsRepositoryMock{
		listMock: []Migration{pending_migration},
	}
	references := &referenceRepositoryMock{}
	var service Service = &ServiceImpl{
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

func TestUp_WhenAllMigrationsAreRan(t *testing.T) {
	migrations := &migrationsRepositoryMock{
		listMock: []Migration{{
			Name: "testing",
			Path: "testing",
		}},
	}
	references := &referenceRepositoryMock{
		listMock: []Reference{{
			Name: "testing",
		}},
	}
	var service Service = &ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	ran, err := service.Up()
	if err != nil {
		t.Error(err)
	}
	expected := Migration{}
	if ran != expected {
		t.Fatal()
	}
}

func TestUp_WhenHasNoMigrations(t *testing.T) {
	migrations := &migrationsRepositoryMock{}
	references := &referenceRepositoryMock{}
	var service Service = &ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	ran, err := service.Up()
	if err != nil {
		t.Error(err)
	}
	expected := Migration{}
	if ran != expected {
		t.Fatal()
	}
}

func TestUp_WhenMigrationsAreCorrupted(t *testing.T) {
	expected_migration := Migration{}
	migrations := &migrationsRepositoryMock{
		listMock: []Migration{
			{
				Name: "testing",
				Path: "testing",
			},
		},
	}
	references := &referenceRepositoryMock{
		listMock: []Reference{{
			Name: "wrong_testing",
		}},
	}
	var service Service = &ServiceImpl{
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

func TestUp_WhenMigrationsAreDisorderly(t *testing.T) {
	migrations := &migrationsRepositoryMock{
		listMock: []Migration{
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
		listMock: []Reference{
			{
				Name: "0_testing",
			},
			{
				Name: "2_testing",
			},
		},
	}
	var service Service = &ServiceImpl{
		Migrations: migrations,
		References: references,
	}
	ran, err := service.Up()
	if err != nil {
		t.Fatal(err)
	}
	expected := Migration{
		Name: "1_testing",
		Path: "testing",
	}
	if ran != expected {
		t.Fatal(ran, " is not ", expected)
	}
}
