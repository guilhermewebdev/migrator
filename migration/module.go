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

func NewMigrationModule(settings conf.Settings, pool lib.DB) (MigrationModule, error) {
	var disk lib.Disk = &lib.DiskImpl{}
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
