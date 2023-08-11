package cli

import (
	"log"

	"github.com/guilhermewebdev/migrator/conf"
	"github.com/guilhermewebdev/migrator/lib"
	"github.com/guilhermewebdev/migrator/migration"
)

func create_migration(pool lib.DB, settings conf.Settings, migration_name string) error {
	module, err := migration.NewMigrationModule(settings, pool)
	response, err := module.Controller().Create(migration_name)
	if err != nil {
		return err
	}
	log.Println(response)
	return nil
}

func up(pool lib.DB, settings conf.Settings) error {
	module, err := migration.NewMigrationModule(settings, pool)
	response, err := module.Controller().Up()
	if err != nil {
		return err
	}
	log.Println(response)
	return nil
}

func unlock(pool lib.DB, settings conf.Settings) error {
	module, err := migration.NewMigrationModule(settings, pool)
	response, err := module.Controller().Unlock()
	if err != nil {
		return err
	}
	log.Println(response)
	return nil
}

func down(pool lib.DB, settings conf.Settings) error {
	module, err := migration.NewMigrationModule(settings, pool)
	response, err := module.Controller().Down()
	if err != nil {
		return err
	}
	log.Println(response)
	return nil
}

func latest(pool lib.DB, settings conf.Settings) error {
	module, err := migration.NewMigrationModule(settings, pool)
	response, err := module.Controller().Latest()
	if err != nil {
		return err
	}
	log.Println(response)
	return nil
}
