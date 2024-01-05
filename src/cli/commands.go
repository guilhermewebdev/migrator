package cli

import (
	"fmt"

	"github.com/guilhermewebdev/migrator/src/conf"
	"github.com/guilhermewebdev/migrator/src/lib"
	"github.com/guilhermewebdev/migrator/src/migration"
	stgs "github.com/guilhermewebdev/migrator/src/settings"
)

func create_migration(pool lib.DB, settings conf.Settings, migration_name string) error {
	migrations := migration.NewMigrationModule(settings, pool)
	response, err := migrations.Create(migration_name)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func up(pool lib.DB, settings conf.Settings) error {
	migrations := migration.NewMigrationModule(settings, pool)
	response, err := migrations.Up()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func unlock(pool lib.DB, settings conf.Settings) error {
	migrations := migration.NewMigrationModule(settings, pool)
	response, err := migrations.Unlock()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func down(pool lib.DB, settings conf.Settings) error {
	migrations := migration.NewMigrationModule(settings, pool)
	response, err := migrations.Down()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func latest(pool lib.DB, settings conf.Settings) error {
	migrations := migration.NewMigrationModule(settings, pool)
	response, err := migrations.Latest()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func settings(settings_file_name string) error {
	settings := stgs.NewSettingsModule()
	response, err := settings.Get(settings_file_name)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
