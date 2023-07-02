package migration

import (
	"bytes"
	"database/sql"
	_ "embed"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/guilhermewebdev/migrator/conf"
	"github.com/guilhermewebdev/migrator/lib"
	"golang.org/x/exp/maps"
)

type ReferenceRepository interface {
	List() ([]Reference, error)
	Up(migration Migration) error
	Down(migration Migration) error
	Lock() error
	Unlock() error
	IsLocked() (bool, error)
	Prepare() error
	GetLast() (Reference, error)
}

type scannable interface {
	Scan(dest ...any) error
}

type ReferenceRepositoryImpl struct {
	DB       lib.DB
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

	//go:embed sql/select_last_reference.sql
	select_last_reference_sql string

	//go:embed sql/delete_reference.sql
	delete_reference_sql string
)

const DB_TIMESTAMP_FORMAT = "2006-01-02 15:04:05"

type P map[string]any

func (r *ReferenceRepositoryImpl) format(query string, values ...P) (string, error) {
	var formatted bytes.Buffer
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
	if err := tmpl.Execute(&formatted, variables); err != nil {
		return "", err
	}
	return formatted.String(), nil
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

func (r *ReferenceRepositoryImpl) query_row(query string, values ...P) (*sql.Row, error) {
	formatted, err := r.format(query, values...)
	if err != nil {
		return &sql.Row{}, err
	}
	row := r.DB.QueryRow(formatted)
	return row, row.Err()
}

func (r *ReferenceRepositoryImpl) scan_ref(row scannable) (Reference, error) {
	var reference Reference
	var date []byte
	err := row.Scan(
		&reference.ID,
		&reference.Name,
		&date,
	)
	if err != nil {
		return reference, err
	}
	without_z := strings.Replace(string(date), "Z", "", 1)
	without_t := strings.Replace(without_z, "T", " ", 1)
	reference.Date, err = time.Parse(DB_TIMESTAMP_FORMAT, without_t)
	return reference, err
}

func (r *ReferenceRepositoryImpl) Prepare() error {
	if _, err := r.exec(create_reference_table_sql); err != nil {
		return err
	}
	if _, err := r.exec(create_lock_table_sql); err != nil {
		return err
	}
	return nil
}

func (r *ReferenceRepositoryImpl) genReferenceId() (int, error) {
	row, err := r.query_row(next_id_sql)
	if err != nil {
		return 0, err
	}
	var next_id int
	row.Scan(&next_id)
	return next_id, row.Err()
}

func (r *ReferenceRepositoryImpl) List() ([]Reference, error) {
	var references []Reference
	query := list_references_sql
	rows, err := r.query(query)
	if err != nil {
		return references, err
	}
	for rows.Next() {
		reference, err := r.scan_ref(rows)
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

func (r *ReferenceRepositoryImpl) Up(m Migration) error {
	if _, err := r.DB.Exec(m.UpQuery); err != nil {
		return err
	}
	id, err := r.genReferenceId()
	if err != nil {
		return err
	}
	if _, err = r.exec(insert_reference_sql, P{
		"id":            id,
		"migration_key": m.Name,
		"created_at":    time.Now().UTC().Format(DB_TIMESTAMP_FORMAT),
	}); err != nil {
		return err
	}
	return nil
}

func (r *ReferenceRepositoryImpl) Down(m Migration) error {
	if _, err := r.DB.Exec(m.DownQuery); err != nil {
		return err
	}
	if _, err := r.exec(delete_reference_sql, P{
		"migration_key": m.Name,
	}); err != nil {
		return err
	}
	return nil
}

func (r *ReferenceRepositoryImpl) IsLocked() (bool, error) {
	row, err := r.query_row(is_locked_sql)
	if err != nil {
		return false, err
	}
	var is_locked bool
	row.Scan(&is_locked)
	return is_locked, row.Err()
}

func (r *ReferenceRepositoryImpl) setLock(is_locked bool) error {
	var exists_data bool
	row, err := r.query_row(check_lock_sql)
	if err != nil {
		return err
	}
	if err := row.Scan(&exists_data); err != nil {
		return err
	}
	var query string
	if exists_data {
		query = update_lock_sql
	} else {
		query = insert_lock_sql
	}
	_, err = r.exec(query, P{"is_locked": is_locked})
	return err
}

func (r *ReferenceRepositoryImpl) Lock() error {
	return r.setLock(true)
}

func (r *ReferenceRepositoryImpl) Unlock() error {
	return r.setLock(false)
}

func (r *ReferenceRepositoryImpl) GetLast() (Reference, error) {
	row, err := r.query_row(select_last_reference_sql)
	if err != nil {
		return Reference{}, err
	}
	reference, err := r.scan_ref(row)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return Reference{}, fmt.Errorf("No migrations to rollback")
	}
	return reference, err
}
