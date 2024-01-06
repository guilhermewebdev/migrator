package settings

import "github.com/guilhermewebdev/migrator/src/lib"

func NewSettingsModule() Controller {
	var disk lib.Disk = &lib.DiskImpl{}
	var repository SettingsRepository = &SettingsRepositoryImpl{
		Disk: disk,
	}
	var service Service = &ServiceImpl{
		Settings: repository,
	}
	var controller = &ControllerImpl{
		Service: service,
	}
	return controller
}
