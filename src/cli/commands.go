package cli

import (
	"fmt"

	"github.com/guilhermewebdev/migrator/src/lib"
	"github.com/guilhermewebdev/migrator/src/migration"
	stgs "github.com/guilhermewebdev/migrator/src/settings"
	"gopkg.in/yaml.v2"
)

func create_migration(settings stgs.Settings, migration_name string) error {
	migrations := migration.NewMigrationModule(settings, nil)
	response, err := migrations.Create(migration_name)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func up(pool lib.DB, settings stgs.Settings) error {
	migrations := migration.NewMigrationModule(settings, pool)
	response, err := migrations.Up()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func unlock(pool lib.DB, settings stgs.Settings) error {
	migrations := migration.NewMigrationModule(settings, pool)
	response, err := migrations.Unlock()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func down(pool lib.DB, settings stgs.Settings) error {
	migrations := migration.NewMigrationModule(settings, pool)
	response, err := migrations.Down()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func latest(pool lib.DB, settings stgs.Settings) error {
	migrations := migration.NewMigrationModule(settings, pool)
	response, err := migrations.Latest()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func settings(settings stgs.Settings) error {
	yamlData, err := yaml.Marshal(&settings)
	if err != nil {
		return err
	}
	fmt.Println(string(yamlData))
	return nil
}
