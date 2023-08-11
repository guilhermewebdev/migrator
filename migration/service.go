package migration

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/guilhermewebdev/migrator/lib"
)

type Service interface {
	Create(name string) error
	Up() (Migration, error)
	Unlock() error
	Down() (Migration, error)
	Latest() ([]Migration, error)
}

type ServiceImpl struct {
	Migrations MigrationRepository
	References ReferenceRepository
}

func (s *ServiceImpl) semaphore() func() {
	if err := s.References.Prepare(); err != nil {
		log.Fatal(err)
	}
	is_locked, err := s.References.IsLocked()
	if err != nil {
		log.Fatal(err)
	}
	if is_locked {
		log.Fatal("The migrations are locked. Unlock it to continue.")
	}
	if err = s.References.Lock(); err != nil {
		log.Fatal(err)
	}
	return func() {
		if err = s.References.Unlock(); err != nil {
			log.Fatal(err)
		}
	}
}

func (s *ServiceImpl) relateMigrationWithReference() ([]Relation, error) {
	references, err := s.References.List()
	if err != nil {
		return []Relation{}, err
	}
	migrations, err := s.Migrations.List()
	if err != nil {
		return []Relation{}, err
	}
	if len(references) > len(migrations) {
		return []Relation{}, fmt.Errorf("Migrations are corrupted")
	}
	if len(migrations) == 0 {
		return []Relation{}, nil
	}
	var relations []Relation
	references_related := 0
	for i, migration := range migrations {
		relation := Relation{
			Migration: migration,
			Reference: Reference{},
		}
		if len(references) > i && references[i].Name == migration.Name {
			relation.Reference = references[i]
			references_related += 1
		} else {
			for _, reference := range references {
				if reference.Name == migration.Name {
					relation.Reference = reference
					references_related += 1
					break
				}
			}
		}
		relations = append(relations, relation)
	}
	if references_related < len(references) {
		return []Relation{}, fmt.Errorf("Migrations are corrupted")
	}
	sort.Slice(relations, func(i, j int) bool {
		return relations[i].Migration.Name > relations[j].Migration.Name
	})
	return relations, nil
}

func (s *ServiceImpl) getNextMigration() (Migration, error) {
	relations, err := s.relateMigrationWithReference()
	if err != nil {
		return Migration{}, err
	}
	for _, relation := range relations {
		if relation.Reference.Name == "" {
			return relation.Migration, nil
		}
	}
	return Migration{}, nil
}

func (s *ServiceImpl) listMissingMigrations() ([]Migration, error) {
	relations, err := s.relateMigrationWithReference()
	if err != nil {
		return []Migration{}, err
	}
	migrations := []Migration{}
	for _, relation := range relations {
		if relation.Reference.Name == "" && relation.Reference.ID == "" {
			migrations = append(migrations, relation.Migration)
		}
	}
	return migrations, nil
}

func (s *ServiceImpl) Create(name string) error {
	snake_case_name := lib.SnakeCase(name)
	now := time.Now().UnixMilli()
	migration_name := fmt.Sprint(now) + "_" + snake_case_name
	return s.Migrations.Create(migration_name)
}

func (s *ServiceImpl) Up() (Migration, error) {
	defer s.semaphore()()
	empty := Migration{}
	migration, err := s.getNextMigration()
	if err != nil {
		return empty, err
	}
	if migration == empty {
		return empty, nil
	}
	err = s.References.Up(migration)
	if err != nil {
		return empty, err
	}
	return migration, nil
}

func (s *ServiceImpl) Unlock() error {
	if err := s.References.Prepare(); err != nil {
		return err
	}
	return s.References.Unlock()
}

func (s *ServiceImpl) Down() (Migration, error) {
	defer s.semaphore()()
	empty := Migration{}
	last_reference, err := s.References.GetLast()
	if err != nil {
		return empty, err
	}
	migration, err := s.Migrations.Read(last_reference.Name)
	if err != nil {
		return empty, err
	}
	err = s.References.Down(migration)
	return migration, err
}

func (s *ServiceImpl) Latest() ([]Migration, error) {
	defer s.semaphore()()
	performed_migrations := []Migration{}
	missing_migrations, err := s.listMissingMigrations()
	if err != nil {
		return performed_migrations, err
	}
	for _, migration := range missing_migrations {
		err = s.References.Up(migration)
		if err != nil {
			return performed_migrations, err
		}
		performed_migrations = append(performed_migrations, migration)
	}
	return performed_migrations, nil
}
