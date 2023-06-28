package migration

import (
	"github.com/guilhermewebdev/migrator/conf"
	"github.com/guilhermewebdev/migrator/lib"
)

type MigrationModule interface {
	Controller() Controller
}

type MigrationModuleImpl struct {
	controller Controller
}

func (mod *MigrationModuleImpl) Controller() Controller {
	return mod.controller
}

func NewMigrationModule(settings conf.Settings) (MigrationModule, error) {
	var disk lib.Disk = &lib.DiskImpl{}
	pool, err := lib.ConnectDB(lib.ConnectionParams{
		DSN:    settings.DBDSN,
		Driver: settings.DBDriver,
	})
	if err != nil {
		return nil, err
	}
	var migrations MigrationRepository = &MigrationRepositoryImpl{
		Disk:     disk,
		Settings: settings,
	}
	var db ReferenceRepository = &ReferenceRepositoryImpl{
		Settings: settings,
		DB:       pool,
	}
	var service Service = &ServiceImpl{
		Migrations: migrations,
		References: db,
	}
	var controller Controller = &ControllerImpl{
		Service: service,
	}
	var module MigrationModule = &MigrationModuleImpl{
		controller: controller,
	}
	return module, nil
}
