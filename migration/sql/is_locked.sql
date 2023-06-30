SELECT COALESCE(is_locked, false) AS is_locked
FROM {{.migrations_lock_table}}
LIMIT 1;