package settings

import "fmt"

type Controller interface {
	Get(file_name string) (Settings, error)
	Init(file_path string) (string, error)
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

func (c *ControllerImpl) Init(file_path string) (string, error) {
	if err := c.Service.Init(file_path); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s conf file was created", file_path), nil
}
