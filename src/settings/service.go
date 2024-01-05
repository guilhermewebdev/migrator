package settings

type SettingsService interface {
	Get(settings_file_name string) (Settings, error)
}

type SettingsServiceImpl struct {
	Settings SettingsRepository
}

func (s *SettingsServiceImpl) get_default_settings() Settings {
	return Settings{
		MigrationsDir:       "./migrations",
		MigrationsTableName: "migrations",
		DB_DSN:              "",
		DB_Driver:           "",
	}
}

func (s *SettingsServiceImpl) combine_settings(stgs ...Settings) Settings {
	var final_settings Settings
	for _, current := range stgs {
		if current.MigrationsDir != "" {
			final_settings.MigrationsDir = current.MigrationsDir
		}
		if current.MigrationsTableName != "" {
			final_settings.MigrationsTableName = current.MigrationsTableName
		}
		if current.DB_DSN != "" {
			final_settings.DB_DSN = current.DB_DSN
		}
		if current.DB_Driver != "" {
			final_settings.DB_Driver = current.DB_Driver
		}
	}
	return final_settings
}

func (s *SettingsServiceImpl) Get(settings_file_name string) (Settings, error) {
	initial := s.get_default_settings()
	env_settings, err := s.Settings.GetFromEnv()
	if err != nil {
		return env_settings, err
	}
	file_settings, err := s.Settings.GetFromFile(settings_file_name)
	if err != nil {
		return file_settings, err
	}
	settings := s.combine_settings(initial, env_settings, file_settings)
	return settings, nil
}
