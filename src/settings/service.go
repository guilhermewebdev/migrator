package settings

type SettingsService interface {
	Get(settings_file_name string) (Settings, error)
}

type SettingsServiceImpl struct {
	Settings SettingsRepository
}

func (s *SettingsServiceImpl) getDefaultSettings() Settings {
	return Settings{
		MigrationsDir:       "./migrations",
		MigrationsTableName: "migrations",
		DB_DSN:              "",
		DB_Driver:           "",
	}
}

func (s *SettingsServiceImpl) Get(settings_file_name string) (Settings, error) {
	initial := s.getDefaultSettings()
	env_settings, err := s.Settings.GetFromEnv(initial)
	if err != nil {
		return env_settings, err
	}
	settings, err := s.Settings.GetFromFile(settings_file_name, env_settings)
	if err != nil {
		return settings, err
	}
	return settings, nil
}
