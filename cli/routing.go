package cli

import (
	"time"

	"github.com/guilhermewebdev/migrator/conf"
	lib_cli "github.com/urfave/cli/v2"
)

func load_settings(cCtx *lib_cli.Context) conf.Settings {
	file_name_from_args := cCtx.StringSlice("conf-file")
	var settings_file string = "migrator.yml"
	if len(file_name_from_args) > 0 {
		settings_file = file_name_from_args[0]
	}
	settings := conf.LoadSettings(settings_file)
	migrations_dir_from_args := cCtx.StringSlice("migrations")
	if len(migrations_dir_from_args) > 0 {
		settings.MigrationsDir = migrations_dir_from_args[0]
	}
	dsn_from_args := cCtx.StringSlice("dsn")
	if len(dsn_from_args) > 0 {
		settings.DBDSN = dsn_from_args[0]
	}
	driver_from_args := cCtx.StringSlice("driver")
	if len(driver_from_args) > 0 {
		settings.DBDriver = driver_from_args[0]
	}
	return settings
}

func BuildRouter() *lib_cli.App {
	app := &lib_cli.App{
		Name:                 "migrate",
		Version:              "0.0.0",
		Compiled:             time.Now(),
		EnableBashCompletion: true,
		Flags: []lib_cli.Flag{
			&lib_cli.StringFlag{
				Name:    "conf-file",
				Aliases: []string{"c"},
				Value:   "migrator.yml",
				Usage:   "Load configuration from `FILE`",
				EnvVars: []string{"CONF_FILE"},
			},
			&lib_cli.StringFlag{
				Name:    "migrations",
				Aliases: []string{"m"},
				Value:   "./migrations",
				Usage:   "Select the migrations directory",
				EnvVars: []string{"MIGRATIONS_DIR"},
			},
			&lib_cli.StringFlag{
				Name:    "dsn",
				Aliases: []string{"d"},
				Usage:   "Database connection string",
				EnvVars: []string{"DB_DSN"},
			},
			&lib_cli.StringFlag{
				Name:    "driver",
				Aliases: []string{"r"},
				Usage:   "Database driver (mysql, postgres...)",
				EnvVars: []string{"DB_DRIVER"},
			},
		},
		Authors: []*lib_cli.Author{
			{
				Name:  "Guilherme Isaías",
				Email: "guilherme@guilhermeweb.dev",
			},
		},
		Commands: []*lib_cli.Command{
			{
				Name:  "new",
				Usage: "Creates a new migration",
				Action: func(cCtx *lib_cli.Context) error {
					settings := load_settings(cCtx)
					return create_migration(settings, cCtx.Args().First())
				},
			},
			{
				Name:  "up",
				Usage: "Execute the next migration",
				Action: func(cCtx *lib_cli.Context) error {
					settings := load_settings(cCtx)
					return up(settings)
				},
			},
		},
	}
	return app
}
