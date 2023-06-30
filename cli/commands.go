package cli

import (
	"database/sql"
	"log"

	"github.com/guilhermewebdev/migrator/conf"
	"github.com/guilhermewebdev/migrator/migration"
)

func create_migration(pool *sql.DB, settings conf.Settings, migration_name string) error {
	module, err := migration.NewMigrationModule(settings, pool)
	response, err := module.Controller().Create(migration_name)
	if err != nil {
		return err
	}
	log.Print(response)
	return nil
}

func up(pool *sql.DB, settings conf.Settings) error {
	module, err := migration.NewMigrationModule(settings, pool)
	response, err := module.Controller().Up()
	if err != nil {
		return err
	}
	log.Print(response)
	return nil
}

func unlock(pool *sql.DB, settings conf.Settings) error {
	module, err := migration.NewMigrationModule(settings, pool)
	response, err := module.Controller().Unlock()
	if err != nil {
		return err
	}
	log.Print(response)
	return nil

}
