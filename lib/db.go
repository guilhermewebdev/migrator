package lib

import (
	"database/sql"
	"flag"
	"log"
)

type ConnectionParams struct {
	DSN    string
	Driver string
}

func ConnectDB(p ConnectionParams) (*sql.DB, error) {
	dsn := flag.String("dsn", p.DSN, "connection data source name")
	flag.Parse()
	pool, err := sql.Open(p.Driver, *dsn)
	if err != nil {
		log.Fatal("Unable to use data source name", err)
	}
	return pool, err
}
