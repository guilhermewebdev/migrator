SELECT *
FROM {{.migrations_table}}
ORDER BY migration_key DESC
LIMIT 1;