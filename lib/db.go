package lib

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/sijms/go-ora"
	_ "modernc.org/sqlite"
)

type ConnectionParams struct {
	DSN    string
	Driver string
}

func ConnectDB(p ConnectionParams) (*sql.DB, error) {
	if p.DSN == "" {
		return &sql.DB{}, fmt.Errorf("Invalid DSN string")
	}
	if p.Driver == "" {
		return &sql.DB{}, fmt.Errorf("Invalid Driver string")
	}
	pool, err := sql.Open(p.Driver, p.DSN)
	return pool, err
}
