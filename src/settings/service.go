package settings

type SettingsService interface {
	Get(settings_file_name string) Settings
}

type SettingsServiceImpl struct{}
