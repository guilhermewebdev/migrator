package migration

import (
	"bytes"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/guilhermewebdev/migrator/conf"
	"golang.org/x/exp/maps"
)

type ReferenceRepository interface {
	List() ([]Reference, error)
	Run(migration Migration) error
	Lock() error
	Unlock() error
	IsLocked() (bool, error)
	Prepare() error
}

type ReferenceRepositoryImpl struct {
	DB       *sql.DB
	Settings conf.Settings
}

var (
	//go:embed sql/list_references.sql
	list_references_sql string

	//go:embed sql/create_reference_table.sql
	create_reference_table_sql string

	//go:embed sql/create_lock_table.sql
	create_lock_table_sql string

	//go:embed sql/insert_reference.sql
	insert_reference_sql string

	//go:embed sql/is_locked.sql
	is_locked_sql string

	//go:embed sql/next_id.sql
	next_id_sql string

	//go:embed sql/check_lock.sql
	check_lock_sql string

	//go:embed sql/insert_lock.sql
	insert_lock_sql string

	//go:embed sql/update_lock.sql
	update_lock_sql string
)

const DB_TIMESTAMP_FORMAT = "2006-01-02 15:04:05"

type P map[string]any

func (r *ReferenceRepositoryImpl) format(query string, values ...P) (string, error) {
	var msg bytes.Buffer
	variables := P{}
	for _, value := range values {
		maps.Copy(variables, value)
	}
	variables["migrations_table"] = r.Settings.MigrationsTableName
	variables["migrations_lock_table"] = r.Settings.MigrationsTableName + "_lock"
	tmpl, err := template.New("query").Parse(query)
	if err != nil {
		return "", err
	}
	if err := tmpl.Execute(&msg, variables); err != nil {
		return "", err
	}
	log.Println(msg.String(), "\n")
	return msg.String(), nil
}

func (r *ReferenceRepositoryImpl) exec(query string, values ...P) (sql.Result, error) {
	formatted, err := r.format(query, values...)
	if err != nil {
		return nil, err
	}
	return r.DB.Exec(formatted)
}

func (r *ReferenceRepositoryImpl) query(query string, values ...P) (*sql.Rows, error) {
	formatted, err := r.format(query, values...)
	if err != nil {
		return &sql.Rows{}, err
	}
	return r.DB.Query(formatted)
}

func (r *ReferenceRepositoryImpl) query_row(query string, values ...P) *sql.Row {
	formatted, err := r.format(query, values...)
	if err != nil {
		return &sql.Row{}
	}
	return r.DB.QueryRow(formatted)
}

func (r *ReferenceRepositoryImpl) Prepare() error {
	_, err := r.exec(create_reference_table_sql)
	_, err = r.exec(create_lock_table_sql)
	return err
}

func (r *ReferenceRepositoryImpl) genReferenceId() int {
	row := r.query_row(next_id_sql)
	var next_id int
	row.Scan(&next_id)
	return next_id
}

func (r *ReferenceRepositoryImpl) List() ([]Reference, error) {
	var references []Reference
	query := list_references_sql
	rows, err := r.query(query)
	if err != nil {
		return references, err
	}
	for rows.Next() {
		var reference Reference
		var date []byte
		if err := rows.Scan(
			&reference.ID,
			&reference.Name,
			&date,
		); err != nil {
			return references, err
		}
		reference.Date, err = time.Parse(DB_TIMESTAMP_FORMAT, string(date))
		if err != nil {
			return references, err
		}
		references = append(references, reference)
	}
	if !rows.NextResultSet() {
		return references, rows.Err()
	}
	return references, nil
}

func (r *ReferenceRepositoryImpl) Run(m Migration) error {
	log.Println(m.UpQuery, "\n")
	if _, err := r.DB.Exec(m.UpQuery); err != nil {
		return err
	}
	_, err := r.exec(insert_reference_sql, P{
		"id":            r.genReferenceId(),
		"migration_key": m.Name,
		"created_at":    time.Now().UTC().Format(DB_TIMESTAMP_FORMAT),
	})
	if err != nil {
		return err
	}
	fmt.Print("OK: ", m.Name)
	return nil
}

func (r *ReferenceRepositoryImpl) IsLocked() (bool, error) {
	row := r.query_row(is_locked_sql)
	var is_locked bool
	row.Scan(&is_locked)
	return is_locked, row.Err()
}

func (r *ReferenceRepositoryImpl) setLock(is_locked bool) error {
	var exists_data bool
	if err := r.query_row(check_lock_sql).Scan(&exists_data); err != nil {
		return err
	}
	var query string
	if exists_data {
		query = update_lock_sql
	} else {
		query = insert_lock_sql
	}
	_, err := r.exec(query, P{"is_locked": is_locked})
	return err
}

func (r *ReferenceRepositoryImpl) Lock() error {
	return r.setLock(true)
}

func (r *ReferenceRepositoryImpl) Unlock() error {
	return r.setLock(false)
}
