package migration_test

import (
	"os"
	"testing"

	"github.com/guilhermewebdev/migrator/modules/migration"
)

func clear() error {
	err := os.RemoveAll("./test")
	return err
}

func setup() func() {
	clear()
	return func() {
		clear()
	}
}

func TestCreateFile(t *testing.T) {
	defer setup()()
	var repo migration.DiskRepository = &migration.DiskRepositoryImpl{}
	err := repo.CreateFile("./test", "file.sql")
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat("./test/file.sql"); err != nil {
		t.Error(err)
	}
}

func TestCreateDeepFile(t *testing.T) {
	defer setup()()
	var repo migration.DiskRepository = &migration.DiskRepositoryImpl{}
	err := repo.CreateFile("./test/testing/test/test/tes", "file.sql")
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat("./test/testing/test/test/tes/file.sql"); err != nil {
		t.Error(err)
	}
}

func TestCreateTheSameFileMultiplesTimes(t *testing.T) {
	defer setup()()
	var repo migration.DiskRepository = &migration.DiskRepositoryImpl{}
	err := repo.CreateFile("./test/testing/test/test/tes", "file.sql")
	repo.CreateFile("./test/testing/test/test/tes", "file.sql")
	repo.CreateFile("./test/testing/test/test/tes", "file.sql")
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat("./test/testing/test/test/tes/file.sql"); err != nil {
		t.Error(err)
	}
}
