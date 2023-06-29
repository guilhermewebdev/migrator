CREATE TABLE IF NOT EXISTS migrations_lock (
    id INT NOT NULL,
    is_locked BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);