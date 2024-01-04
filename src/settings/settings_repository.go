package settings

type SettingsRepository interface {
	GetFromEnv() (Settings, error)
	GetFromFile(file_name string) (Settings, error)
}
