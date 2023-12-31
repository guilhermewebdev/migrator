package settings

type SettingsRepository interface {
	GetFromEnv(default_settings Settings) (Settings, error)
	GetFromFile(file_name string, default_settings Settings) (Settings, error)
}
