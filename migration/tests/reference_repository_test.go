package migration_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/guilhermewebdev/migrator/migration"
)

func TestList(t *testing.T) {
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
	repo := migration.ReferenceRepositoryImpl{
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
