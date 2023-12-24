package lib

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type Disk interface {
	Create(path_name string, file_name string) error
	List(dir string) ([]string, error)
	Read(file_path string) (string, error)
}

type DiskImpl struct{}

func (d *DiskImpl) Create(path_name string, file_name string) error {
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

func (d *DiskImpl) List(dir string) ([]string, error) {
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

func (d *DiskImpl) Read(file_path string) (string, error) {
	full_file_name, _ := filepath.Abs(file_path)
	data, err := os.ReadFile(full_file_name)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
