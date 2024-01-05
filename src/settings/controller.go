package settings

import "gopkg.in/yaml.v2"

type Controller interface {
	Get(file_name string) (string, error)
}

type ControllerImpl struct {
	Service Service
}

func (c *ControllerImpl) Get(file_name string) (string, error) {
	settings, err := c.Service.Get(file_name)
	if err != nil {
		return "Failed to load settings", err
	}
	yamlData, err := yaml.Marshal(&settings)
	return string(yamlData), err
}
