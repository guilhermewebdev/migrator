package settings

type Controller interface {
	Get(file_name string) (Settings, error)
}

type ControllerImpl struct {
	Service Service
}

func (c *ControllerImpl) Get(file_name string) (Settings, error) {
	settings, err := c.Service.Get(file_name)
	if err != nil {
		return Settings{}, err
	}
	return settings, err
}
