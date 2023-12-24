INSERT 
    INTO {{.migrations_table}} (id, migration_key, created_at) 
    VALUES ({{.id}}, '{{.migration_key}}', '{{.created_at}}');