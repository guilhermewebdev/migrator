package migration

import (
	"github.com/guilhermewebdev/migrator/src/lib"
	stgs "github.com/guilhermewebdev/migrator/src/settings"
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

func NewMigrationModule(settings stgs.Settings, pool lib.DB) Controller {
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
	return controller
}
