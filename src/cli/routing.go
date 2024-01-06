package cli

import (
	"log"
	"time"

	"github.com/guilhermewebdev/migrator/src/lib"
	stgs "github.com/guilhermewebdev/migrator/src/settings"
	lib_cli "github.com/urfave/cli/v2"
)

type context struct {
	s    stgs.Settings
	c    *lib_cli.Context
	pool lib.DB
}

func get_settings_file_name(ctx *lib_cli.Context) string {
	file_name_from_args := ctx.String("conf-file")
	var settings_file string = "./migrator.yml"
	if len(file_name_from_args) > 0 {
		settings_file = file_name_from_args
	}
	return settings_file
}

func load_settings(ctx *lib_cli.Context) (stgs.Settings, error) {
	settings_file := get_settings_file_name(ctx)
	settings, err := stgs.NewSettingsModule().Get(settings_file)
	if err != nil {
		return stgs.Settings{}, err
	}
	migrations_dir_from_args := ctx.String("migrations")
	if len(migrations_dir_from_args) > 0 {
		settings.MigrationsDir = migrations_dir_from_args
	}
	dsn_from_args := ctx.String("dsn")
	if len(dsn_from_args) > 0 {
		settings.DB_DSN = dsn_from_args
	}
	driver_from_args := ctx.String("driver")
	if len(driver_from_args) > 0 {
		settings.DB_Driver = driver_from_args
	}
	table_name_from_args := ctx.String("table")
	if len(table_name_from_args) > 0 {
		settings.MigrationsTableName = table_name_from_args
	}
	return settings, nil
}

func call(action func(context) error) lib_cli.ActionFunc {
	return func(ctx *lib_cli.Context) error {
		settings, err := load_settings(ctx)
		if err != nil {
			return err
		}
		return action(context{settings, ctx, nil})
	}
}

func db(action func(context) error) lib_cli.ActionFunc {
	return func(ctx *lib_cli.Context) error {
		settings, err := load_settings(ctx)
		if err != nil {
			return err
		}
		pool, err := lib.ConnectDB(lib.ConnectionParams{
			DSN:    settings.DB_DSN,
			Driver: settings.DB_Driver,
		})
		if err != nil {
			return err
		}
		defer func() {
			err := pool.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
		return action(context{settings, ctx, pool})
	}
}

func BuildRouter() *lib_cli.App {
	app := &lib_cli.App{
		Name:                 "Migrator",
		Usage:                "Manage your databases with migrations",
		Version:              "0.2",
		Compiled:             time.Now().UTC(),
		EnableBashCompletion: true,
		HelpName:             "migrate",
		DefaultCommand:       "help",
		Flags: []lib_cli.Flag{
			&lib_cli.StringFlag{
				Name:    "conf-file",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE` (default: \"migrator.yml\")",
				EnvVars: []string{"CONF_FILE"},
			},
			&lib_cli.StringFlag{
				Name:    "migrations",
				Aliases: []string{"m"},
				Usage:   "Select the migrations directory (default: \"./migrations\")",
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
				Usage:   "Database driver (mysql, postgres, sqlserver, sqlite, sqlite3 or oracle)",
				EnvVars: []string{"DB_DRIVER"},
			},
			&lib_cli.StringFlag{
				Name:    "table",
				Aliases: []string{"t"},
				Usage:   "Migrations table name(default: \"migrations\") ",
				EnvVars: []string{"MIGRATIONS_TABLE"},
			},
		},
		Authors: []*lib_cli.Author{
			{
				Name:  "Guilherme Isaías",
				Email: "guilherme@cibernetica.dev",
			},
		},
		Commands: []*lib_cli.Command{
			{
				Name:  "init",
				Usage: "Create a new config file",
				Action: func(ctx *lib_cli.Context) error {
					return init_command(get_settings_file_name(ctx))
				},
			},
			{
				Name:  "new",
				Usage: "Creates a new migration",
				Action: call(func(ctx context) error {
					return create_migration_command(ctx.s, ctx.c.Args().First())
				}),
			},
			{
				Name:  "up",
				Usage: "Execute the next migration",
				Action: db(func(ctx context) error {
					return up_command(ctx.pool, ctx.s)
				}),
			},
			{
				Name:  "down",
				Usage: "Rollback the last migration",
				Action: db(func(ctx context) error {
					return down_command(ctx.pool, ctx.s)
				}),
			},
			{
				Name:  "unlock",
				Usage: "Unlock migrations",
				Action: db(func(ctx context) error {
					return unlock_command(ctx.pool, ctx.s)
				}),
			},
			{
				Name:  "latest",
				Usage: "Perform missing migrations",
				Action: db(func(ctx context) error {
					return latest_command(ctx.pool, ctx.s)
				}),
			},
			{
				Name:  "settings",
				Usage: "Show settings",
				Action: call(func(ctx context) error {
					return settings_command(ctx.s)
				}),
			},
		},
	}
	return app
}
