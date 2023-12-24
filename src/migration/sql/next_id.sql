SELECT MAX(id) + 1 AS next_id
FROM {{.migrations_table}};