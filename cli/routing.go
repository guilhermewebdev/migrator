package cli

import (
	"time"

	lib_cli "github.com/urfave/cli/v2"
)

func BuildRouter() *lib_cli.App {
	app := &lib_cli.App{
		Name:     "migrate",
		Version:  "0.0.0",
		Compiled: time.Now(),
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
					return create_migration(cCtx.Args().First())
				},
			},
			{
				Name:  "up",
				Usage: "Execute the next migration",
				Action: func(cCtx *lib_cli.Context) error {
					return up()
				},
			},
		},
	}
	return app
}
