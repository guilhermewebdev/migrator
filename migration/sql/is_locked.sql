SELECT COALESCE(is_locked, false) AS is_locked
FROM migrations_lock
LIMIT 1;