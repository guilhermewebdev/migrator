package settings

func NewSettingsModule() Controller {
	var repository SettingsRepository = &SettingsRepositoryImpl{}
	var service Service = &ServiceImpl{
		Settings: repository,
	}
	var controller = &ControllerImpl{
		Service: service,
	}
	return controller
}
