INSERT INTO {{.migrations_lock_table}} (id, is_locked)
SELECT 1, {{.is_locked}};
