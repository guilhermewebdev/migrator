package settings

type Settings struct {
	MigrationsDir       string `yaml:"migrations_dir"`
	MigrationsTableName string `yaml:"migrations_table_name"`
	DB_DSN              string `yaml:"db_dsn"`
	DB_Driver           string `yaml:"db_driver"`
}
