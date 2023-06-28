package cli

import (
	"log"

	"github.com/guilhermewebdev/migrator/conf"
	"github.com/guilhermewebdev/migrator/migration"
)

func create_migration(settings conf.Settings, migration_name string) error {
	module, err := migration.NewMigrationModule(settings)
	response, err := module.Controller().Create(migration_name)
	if err != nil {
		return err
	}
	log.Print(response)
	return nil
}

func up(settings conf.Settings) error {
	module, err := migration.NewMigrationModule(settings)
	response, err := module.Controller().Up()
	if err != nil {
		return err
	}
	log.Print(response)
	return nil
}
