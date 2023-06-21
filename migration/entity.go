package migration

import "time"

type Migration struct {
	ID        string
	Name      string
	UpQuery   string
	DownQuery string
}

type Reference struct {
	ID    string
	Name  string
	Date  time.Time
	Order int
}

type Settings struct {
	MigrationsDir string
}
