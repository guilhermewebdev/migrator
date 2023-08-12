SELECT COUNT(id) AS lock_count
FROM {{.migrations_lock_table}};