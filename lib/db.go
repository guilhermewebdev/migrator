package lib

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora"
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
