package migration

type DiskRepository interface {
	CreateFile(path string, name string) (bool, error)
	List(dir string) ([]string, error)
	Read(path string) (string, error)
}
