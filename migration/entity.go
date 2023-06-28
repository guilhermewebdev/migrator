package migration

import "time"

type Migration struct {
	Name      string
	Path      string
	UpQuery   string
	DownQuery string
}

type Reference struct {
	ID    string
	Name  string
	Date  time.Time
	Order int
}

type Relation struct {
	Migration *Migration
	Reference *Reference
}
