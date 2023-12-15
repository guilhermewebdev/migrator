# Migrator - Database Migration Command Line Tool

## Overview

**Migrator** is a command-line tool designed to simplify the management of databases using migrations. It enables users to create, execute, and roll back migrations, ensuring a smooth and controlled evolution of the database schema.

## Version

Current version: 0.1

## Author

Guilherme Isaías <guilherme@guilhermeweb.dev>

## Usage

```shell
migrate [global options] command [command options] [arguments...]
```

## Commands

1. **new**
   - Creates a new migration.

2. **up**
   - Executes the next migration.

3. **down**
   - Rolls back the last migration.

4. **unlock**
   - Unlocks migrations.

5. **latest**
   - Performs missing migrations.

6. **help, h**
   - Shows a list of commands or help for one command.

## Global Options

- `--conf-file FILE, -c FILE`
  - Load configuration from FILE (default: "migrator.yml").

- `--migrations value, -m value`
  - Select the migrations directory (default: "./migrations").

- `--dsn value, -d value`
  - Database connection string.

- `--driver value, -r value`
  - Database driver (mysql, postgres, sqlserver, sqlite, sqlite3, or oracle).

- `--table value, -t value`
  - Migrations table name (default: "migrations").

- `--help, -h`
  - Show help.

- `--version, -v`
  - Print the version.

## Example

```shell
# Create a new migration
migrate new -n migration_name

# Execute the next migration
migrate up

# Rollback the last migration
migrate down

# Unlock migrations
migrate unlock

# Perform missing migrations
migrate latest
```

## Configuration File

By default, Migrate looks for a configuration file named "migrator.yml" for global settings. You can specify an alternative configuration file using the `--conf-file` option.

Feel free to customize the configuration file to match your project's requirements.

---

For additional information on each command and its options, use:

```shell
migrate [command] --help
```

Thank you for using Migrate! If you have any questions or feedback, please contact Guilherme Isaías at <guilherme@guilhermeweb.dev>.