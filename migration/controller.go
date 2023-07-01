package migration

import "fmt"

type Controller interface {
	Create(name string) (string, error)
	Up() (string, error)
	Unlock() (string, error)
	Down() (string, error)
}

type ControllerImpl struct {
	Service Service
}

func (c *ControllerImpl) Create(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("You should inform the migration name")
	}
	if len(name) > 100 {
		return "", fmt.Errorf("Migration name must not exceed 100 characters")
	}
	if err := c.Service.Create(name); err != nil {
		return "Failed to create the migration", err
	}
	return "Migration was created successfully", nil
}

func (c *ControllerImpl) Up() (string, error) {
	migration, err := c.Service.Up()
	if err != nil {
		return "Failed to execute migration", err
	}
	empty := Migration{}
	if migration == empty {
		return "No migrations to apply", nil
	}
	return "The migration \"" + migration.Name + "\" was ran.", nil
}

func (c *ControllerImpl) Unlock() (string, error) {
	if err := c.Service.Unlock(); err != nil {
		return "Failed to unlock migrations", err
	}
	return "The migrations are unlocked", nil
}

func (c *ControllerImpl) Down() (string, error) {
	migration, err := c.Service.Down()
	if err != nil {
		return "Failed to execute migration", err
	}
	empty := Migration{}
	if migration == empty {
		return "No migrations to apply", nil
	}
	return "The migration \"" + migration.Name + "\" was rolled back.", nil
}
