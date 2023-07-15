package migration_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mod "github.com/guilhermewebdev/migrator/migration"
)

func TestReferenceRepository_List(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mocked_rows := sqlmock.NewRows([]string{
		"ID", "migration_key", "created_at",
	}).AddRow(2, "migration_test", "2006-01-02 15:04:05")
	mock.ExpectQuery("^SELECT \\* FROM migrations").
		WillReturnRows(mocked_rows)
	repo := mod.ReferenceRepositoryImpl{
		Settings: get_settings(),
		DB:       db,
	}
	list, err := repo.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 {
		t.Fatal(list)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestReferenceRepository_Up(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("^CREATE").WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectQuery("^SELECT MAX").
		WillReturnRows(sqlmock.NewRows([]string{"next_id"}).AddRow(2))
	mock.ExpectExec("^INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
	migration := mod.Migration{
		UpQuery: "CREATE TABLE test;",
	}
	repo := mod.ReferenceRepositoryImpl{
		Settings: get_settings(),
		DB:       db,
	}
	err = repo.Up(migration)
	if err != nil {
		t.Fatal(err)
	}
}
