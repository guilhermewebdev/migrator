package migration

type Controller interface {
	Create(name string) (string, error)
}

type ControllerImpl struct {
	Service Service
}

func (controller *ControllerImpl) Create(name string) (string, error) {
	if err := controller.Service.Create(name); err != nil {
		return "Failed to create the migration", err
	}
	return "Migration was created successfully", nil
}
