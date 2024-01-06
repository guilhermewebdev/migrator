package lib_test

import (
	"os"
	"testing"

	"github.com/guilhermewebdev/migrator/src/lib"
)

func clear() {
	os.RemoveAll("./test")
	os.RemoveAll("./mocks")
}

func createMocks() {
	var repo lib.Disk = &lib.DiskImpl{}
	repo.Create("./mocks/file.test")
	os.WriteFile("./mocks/file.test", []byte("hello"), 0644)
}

func setup() func() {
	clear()
	createMocks()
	return func() {
		clear()
	}
}

func TestCreateFile(t *testing.T) {
	defer setup()()
	var repo lib.Disk = &lib.DiskImpl{}
	err := repo.Create("./test/file.sql")
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat("./test/file.sql"); err != nil {
		t.Error(err)
	}
}

func TestCreateDeepFile(t *testing.T) {
	defer setup()()
	var repo lib.Disk = &lib.DiskImpl{}
	err := repo.Create("./test/testing/test/test/tes/file.sql")
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat("./test/testing/test/test/tes/file.sql"); err != nil {
		t.Error(err)
	}
}

func TestCreateTheSameFileMultiplesTimes(t *testing.T) {
	defer setup()()
	var repo lib.Disk = &lib.DiskImpl{}
	err := repo.Create("./test/testing/test/test/tes/file.sql")
	repo.Create("./test/testing/test/test/tes/file.sql")
	repo.Create("./test/testing/test/test/tes/file.sql")
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat("./test/testing/test/test/tes/file.sql"); err != nil {
		t.Error(err)
	}
}

func TestListDirectories(t *testing.T) {
	defer setup()()
	var repo lib.Disk = &lib.DiskImpl{}
	names, err := repo.List("./mocks")
	if err != nil {
		t.Error(err)
	}
	if len(names) < 1 {
		t.Fail()
	}
	var this_file_was_found bool
	for _, name := range names {
		if name == "file.test" {
			this_file_was_found = true
		}
	}
	if !this_file_was_found {
		t.Fail()
	}
}

func TestReadFile(t *testing.T) {
	defer setup()()
	var repo lib.Disk = &lib.DiskImpl{}
	data, err := repo.Read("./mocks/file.test")
	if err != nil {
		t.Error(err)
	}
	if data != "hello" {
		t.Fail()
	}
}
