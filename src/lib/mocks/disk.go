package lib_mocks

import "path"

type DiskMock struct {
	Creations                               []string
	Lists                                   []string
	Reads                                   []string
	CreationMock                            error
	ListMock                                []string
	ListErrorMock                           error
	ReadMock                                string
	ReadErrorMock                           error
	SearchFileInParentDirectories_file_name string
	SearchFileInParentDirectories_return    string
	SearchFileInParentDirectories_error     error
}

func (disk *DiskMock) Create(path_name string, file_name string) error {
	full_path := path.Join(path_name, file_name)
	disk.Creations = append(disk.Creations, full_path)
	return disk.CreationMock
}

func (disk *DiskMock) List(dir string) ([]string, error) {
	disk.Lists = append(disk.Lists, dir)
	return disk.ListMock, disk.ListErrorMock
}

func (disk *DiskMock) Read(file_path string) (string, error) {
	disk.Reads = append(disk.Reads, file_path)
	return disk.ReadMock, disk.ReadErrorMock
}

func (disk *DiskMock) SearchFileInParentDirectories(file_name string) (string, error) {
	disk.SearchFileInParentDirectories_file_name = file_name
	return disk.SearchFileInParentDirectories_return, disk.SearchFileInParentDirectories_error
}
