package migration

import (
	"fmt"
	"sort"
	"time"

	"github.com/guilhermewebdev/migrator/lib"
)

type Service interface {
	Create(name string) error
	Up() (Migration, error)
}

type ServiceImpl struct {
	Migrations MigrationRepository
	References ReferenceRepository
}

func (s *ServiceImpl) Create(name string) error {
	snake_case_name := lib.SnakeCase(name)
	now := time.Now().UnixMilli()
	migration_name := fmt.Sprint(now) + "_" + snake_case_name
	return s.Migrations.Create(migration_name)
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
		return []Relation{}, fmt.Errorf("No migrations to apply.")
	}
	var relations []Relation
	references_related := 0
	for i, migration := range migrations {
		relation := Relation{
			Migration: &migration,
		}
		if references[i].Name == migration.Name {
			relation.Reference = &references[i]
			references_related += 1
		} else {
			for _, reference := range references {
				if reference.Name == migration.Name {
					relation.Reference = &references[i]
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
			return *relation.Migration, nil
		}
	}
	return Migration{}, fmt.Errorf("No migrations to apply")
}

func (s *ServiceImpl) Up() (Migration, error) {
	migration, err := s.getNextMigration()
	if err != nil {
		return Migration{}, err
	}
	err = s.References.Run(migration)
	if err != nil {
		return Migration{}, err
	}
	return migration, nil
}
