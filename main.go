package main

import (
	"log"
	"os"
	"time"

	"github.com/guilhermewebdev/migrator/cli"
)

func main() {
	os.Setenv("TZ", "UTC")
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc
	app := cli.BuildRouter()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
