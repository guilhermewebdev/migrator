package migration

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type DiskRepository interface {
	CreateFile(path_name string, file_name string) error
	List(dir string) ([]string, error)
	// Read(path string) (string, error)
}

type DiskRepositoryImpl struct{}

func (repo *DiskRepositoryImpl) CreateFile(path_name string, file_name string) error {
	directory, _ := filepath.Abs(path_name)
	if _, err := os.Stat(path_name); err != nil {
		if err := os.MkdirAll(directory, fs.ModePerm); err != nil {
			return err
		}
	}
	full_file_name := path.Join(directory, file_name)
	file, err := os.Create(full_file_name)
	defer file.Close()
	if err != nil {
		return err
	}
	file.Chmod(0644)
	return nil
}

func (repo *DiskRepositoryImpl) List(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, nil
	}
	names := []string{}
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names, nil
}
