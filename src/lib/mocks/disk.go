package lib_mocks

type DiskMock struct {
	Creations                               []string
	Lists                                   []string
	Reads                                   []string
	Writes                                  [][]string
	CreationMock                            error
	ListMock                                []string
	ListErrorMock                           error
	ReadMock                                string
	ReadErrorMock                           error
	SearchFileInParentDirectories_file_name string
	SearchFileInParentDirectories_return    string
	SearchFileInParentDirectories_error     error
	WriteError                              error
}

func (disk *DiskMock) Create(file_path string) error {
	disk.Creations = append(disk.Creations, file_path)
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

func (disk *DiskMock) Write(file_name string, content string) error {
	writing := []string{file_name, content}
	disk.Writes = append(disk.Writes, writing)
	return disk.WriteError
}
