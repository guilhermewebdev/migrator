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
	//go:embed sql/list_references.sql
	list_references_sql string

	//go:embed sql/create_reference_table.sql
	create_reference_table_sql string

	//go:embed sql/create_lock_table.sql
	create_lock_table_sql string

	//go:embed sql/insert_reference.sql
	insert_reference_sql string

	//go:embed sql/set_lock.sql
	set_lock_sql string

	//go:embed sql/is_locked.sql
	is_locked_sql string

	//go:embed sql/next_id.sql
	next_id_sql string
)

func (r *ReferenceRepositoryImpl) preCreateTables() error {
	_, err := r.DB.Exec(create_reference_table_sql)
	_, err = r.DB.Exec(create_lock_table_sql)
	return err
}

func (r *ReferenceRepositoryImpl) genReferenceId() int {
	row := r.DB.QueryRow(next_id_sql)
	var next_id int
	row.Scan(&next_id)
	return next_id
}

func (r *ReferenceRepositoryImpl) List() ([]Reference, error) {
	var references []Reference
	if err := r.preCreateTables(); err != nil {
		return references, err
	}
	query := list_references_sql
	rows, err := r.DB.Query(query)
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
		insert_reference_sql,
		r.genReferenceId(),
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
	if err := r.preCreateTables(); err != nil {
		return false, err
	}
	query := is_locked_sql
	row := r.DB.QueryRow(
		query,
	)
	var is_locked bool
	row.Scan(&is_locked)
	return is_locked, row.Err()
}

func (r *ReferenceRepositoryImpl) setLock(val bool) error {
	if err := r.preCreateTables(); err != nil {
		return err
	}
	query := set_lock_sql
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
