package migration

import (
	"path"

	"github.com/guilhermewebdev/migrator/conf"
	"github.com/guilhermewebdev/migrator/lib"
)

type MigrationRepository interface {
	Create(name string) error
	List() ([]Migration, error)
	Read(key string) (Migration, error)
}

type MigrationRepositoryImpl struct {
	Disk     lib.Disk
	Settings conf.Settings
}

func (r *MigrationRepositoryImpl) Create(name string) error {
	new_migration_path := path.Join(r.Settings.MigrationsDir, name)
	if err := r.Disk.Create(new_migration_path, "up.sql"); err != nil {
		return err
	}
	if err := r.Disk.Create(new_migration_path, "down.sql"); err != nil {
		return err
	}
	return nil
}

func (r *MigrationRepositoryImpl) List() ([]Migration, error) {
	keys, err := r.Disk.List(r.Settings.MigrationsDir)
	if err != nil {
		return nil, err
	}
	var migrations []Migration
	for _, key := range keys {
		migration, err := r.Read(key)
		if err != nil {
			return migrations, err
		}
		migrations = append(migrations, migration)
	}
	return migrations, nil
}

func (r *MigrationRepositoryImpl) Read(key string) (Migration, error) {
	empty := Migration{}
	dir := path.Join(r.Settings.MigrationsDir, key)
	up, err := r.Disk.Read(path.Join(dir, "up.sql"))
	if err != nil {
		return empty, err
	}
	down, err := r.Disk.Read(path.Join(dir, "down.sql"))
	if err != nil {
		return empty, err
	}
	return Migration{
		Name:      key,
		Path:      path.Join(r.Settings.MigrationsDir, key),
		UpQuery:   up,
		DownQuery: down,
	}, nil
}
