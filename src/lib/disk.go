package lib

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Disk interface {
	Create(file_path string) error
	List(dir string) ([]string, error)
	Read(file_path string) (string, error)
	SearchFileInParentDirectories(file_name string) (string, error)
	Write(file_name string, content string) error
}

type DiskImpl struct{}

func (d *DiskImpl) Create(file_path string) error {
	path_name := filepath.Dir(file_path)
	directory, _ := filepath.Abs(path_name)
	if _, err := os.Stat(path_name); err != nil {
		if err := os.MkdirAll(directory, fs.ModePerm); err != nil {
			return err
		}
	}
	full_file_name, err := filepath.Abs(file_path)
	if err != nil {
		return err
	}
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

func (r *DiskImpl) SearchFileInParentDirectories(file_name string) (string, error) {
	current_dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		file_path := filepath.Join(current_dir, file_name)
		_, err := os.Stat(file_path)
		if err == nil {
			return file_path, nil
		}
		if current_dir == filepath.Dir(current_dir) {
			break
		}
		current_dir = filepath.Dir(current_dir)
	}
	return "", nil
}

func (r *DiskImpl) Write(file_path string, content string) error {
	return os.WriteFile(file_path, []byte(content), 0644)
}
