package cli

import (
	"log"

	"github.com/guilhermewebdev/migrator/migration"
)

func create_migration(migration_name string) error {
	module, err := migration.NewMigrationModule()
	response, err := module.Controller().Create(migration_name)
	if err != nil {
		return err
	}
	log.Print(response)
	return nil
}

func up() error {
	module, err := migration.NewMigrationModule()
	response, err := module.Controller().Up()
	if err != nil {
		return err
	}
	log.Print(response)
	return nil
}
