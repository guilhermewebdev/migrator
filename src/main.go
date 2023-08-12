package main

import (
	"log"
	"os"

	"github.com/guilhermewebdev/migrator/src/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
