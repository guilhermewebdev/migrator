CREATE TABLE IF NOT EXISTS {{.migrations_lock_table}} (
    id INT NOT NULL,
    is_locked BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);