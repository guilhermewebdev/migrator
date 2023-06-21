package migration

import (
	"fmt"
	"path"
	"time"

	"github.com/guilhermewebdev/migrator/lib"
)

type Service interface {
	Create(name string) error
}

type ServiceImpl struct {
	Disk     DiskRepository
	Settings Settings
}

func (service *ServiceImpl) Create(name string) error {
	snake_case_name := lib.SnackCase(name)
	now := time.Now().UnixMilli()
	migration_name := fmt.Sprint(now) + "_" + snake_case_name
	new_migration_path := path.Join(service.Settings.MigrationsDir, migration_name)
	if err := service.Disk.Create(new_migration_path, "up.sql"); err != nil {
		return err
	}
	if err := service.Disk.Create(new_migration_path, "down.sql"); err != nil {
		return err
	}
	return nil
}
