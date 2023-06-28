package migration

import (
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	"github.com/guilhermewebdev/migrator/conf"
)

type ReferenceRepository interface {
	List() ([]Reference, error)
	Run(migration Migration) error
	Lock() error
	Unlock() error
	IsLocked() (bool, error)
}

type ReferenceRepositoryImpl struct {
	DB       *sql.DB
	Settings conf.Settings
}

var (
	// go:embed sql/list_references.sql
	list_references string
	// go:embed sql/create_reference_table.sql
	create_reference_table string
	// go:embed sql/insert_reference.sql
	insert_reference string
	// go:embed sql/set_lock.sql
	set_lock string
	// go:embed sql/is_locked.sql
	is_locked string
)

func (r *ReferenceRepositoryImpl) List() ([]Reference, error) {
	var references []Reference
	query := create_reference_table + list_references
	table_name := r.Settings.MigrationsTableName
	lock_table_name := table_name + "_lock"
	rows, err := r.DB.Query(query, table_name, lock_table_name)
	if err != nil {
		return references, err
	}
	for rows.Next() {
		var reference Reference
		if err := rows.Scan(
			&reference.ID,
			&reference.Name,
			&reference.Date,
		); err != nil {
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
	_, err := r.DB.Exec(m.UpQuery)
	if err != nil {
		return nil
	}
	_, err = r.DB.Exec(
		insert_reference,
		r.Settings.MigrationsTableName,
		m.Name,
		time.Now(),
	)
	if err != nil {
		return nil
	}
	fmt.Print("OK: ", m.Name)
	return nil
}

func (r *ReferenceRepositoryImpl) IsLocked() (bool, error) {
	query := create_reference_table + is_locked
	row := r.DB.QueryRow(
		query,
	)
	var is_locked bool
	row.Scan(&is_locked)
	return is_locked, row.Err()
}

func (r *ReferenceRepositoryImpl) setLock(val bool) error {
	query := create_reference_table + set_lock
	_, err := r.DB.Exec(
		query,
		true,
	)
	return err
}

func (r *ReferenceRepositoryImpl) Lock() error {
	return r.setLock(true)
}

func (r *ReferenceRepositoryImpl) Unlock() error {
	return r.setLock(false)
}
