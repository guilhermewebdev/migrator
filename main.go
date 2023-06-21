package main

import (
	"log"
	"os"

	"github.com/guilhermewebdev/migrator/cli"
)

func main() {
	app := cli.BuildRouter()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
