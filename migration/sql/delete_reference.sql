DELETE
FROM {{.migrations_table}}
WHERE migration_key = {{.migration_key}};