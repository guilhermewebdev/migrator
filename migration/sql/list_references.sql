SELECT *
FROM {{.migrations_table}}
ORDER BY created_at ASC;