package cli

import (
	"database/sql"
	"log"
	"time"

	"github.com/guilhermewebdev/migrator/conf"
	"github.com/guilhermewebdev/migrator/lib"
	lib_cli "github.com/urfave/cli/v2"
)

type context struct {
	s    conf.Settings
	c    *lib_cli.Context
	pool *sql.DB
}

func load_settings(ctx *lib_cli.Context) conf.Settings {
	file_name_from_args := ctx.String("conf-file")
	var settings_file string = "migrator.yml"
	if len(file_name_from_args) > 0 {
		settings_file = file_name_from_args
	}
	settings := conf.LoadSettings(settings_file)
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
	return settings
}

func call(action func(context) error) lib_cli.ActionFunc {
	return func(ctx *lib_cli.Context) error {
		settings := load_settings(ctx)
		pool, err := lib.ConnectDB(lib.ConnectionParams{
			DSN:    settings.DB_DSN,
			Driver: settings.DB_Driver,
		})
		defer func() {
			err := pool.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
		if err != nil {
			return err
		}
		return action(context{settings, ctx, pool})
	}
}

func BuildRouter() *lib_cli.App {
	app := &lib_cli.App{
		Name:                 "migrate",
		Version:              "0.0.0",
		Compiled:             time.Now().UTC(),
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
				Usage:   "Database driver (mysql, postgres, sqlserver, sqlite, sqlite3 or oracle)",
				EnvVars: []string{"DB_DRIVER"},
			},
			&lib_cli.StringFlag{
				Name:    "table",
				Value:   "migrations",
				Aliases: []string{"t"},
				Usage:   "Migrations table name",
				EnvVars: []string{"MIGRATIONS_TABLE"},
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
				Action: call(func(ctx context) error {
					return create_migration(ctx.pool, ctx.s, ctx.c.Args().First())
				}),
			},
			{
				Name:  "up",
				Usage: "Execute the next migration",
				Action: call(func(ctx context) error {
					return up(ctx.pool, ctx.s)
				}),
			},
			{
				Name:  "down",
				Usage: "Rollback the last migration",
				Action: call(func(ctx context) error {
					return down(ctx.pool, ctx.s)
				}),
			},
			{
				Name:  "unlock",
				Usage: "Unlock migrations",
				Action: call(func(ctx context) error {
					return unlock(ctx.pool, ctx.s)
				}),
			},
		},
	}
	return app
}
