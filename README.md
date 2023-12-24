# Migrator - Database Migration Command Line Tool

## Overview

**Migrator** is a command-line tool designed to simplify the management of databases using migrations. It enables users to create, execute, and roll back migrations, ensuring a smooth and controlled evolution of the database schema.

## Version

Current version: 0.1

## Author

Guilherme Isaías <guilherme@guilhermeweb.dev>

## Installation

### Building from source code

#### Prerequisites

Make sure your system meets the following requirements:

Go: Ensure that Go is installed on your machine. You can download and install it from [the official Go website](https://go.dev/doc/install).

Docker (optional): If you plan to build and run the migrator tool in a Docker container, make sure Docker is installed on your system. You can find instructions for installing Docker [here](https://docs.docker.com/engine/install/).

#### Building and Installing
1. Clone the repository to your local machine:

```bash
git clone https://github.com/guilhermewebdev/migrator.git
```

Change into the project directory:

```bash
cd migrator
```

Run the build command using the provided 
Makefile:

```bash
make build
```

This will compile the migrator tool.
Install the migrator tool globally on your system:

```bash
make install
```

This will copy the compiled executable to /usr/local/bin/, making it accessible system-wide.

## Docker

You can also use a pre-built Docker image for Migrator available on Docker Hub. The image is hosted at [guilhermewebdev/migrator](https://hub.docker.com/r/guilhermewebdev/migrator). To run Migrator using Docker, use the following command:

```bash
docker run -it guilhermewebdev/migrator:latest migrate [global options] command [command options] [arguments...]
```

Replace `[global options]`, `command`, `[command options]`, and `[arguments...]` with the specific options and commands you want to execute.

### Example using Docker

```bash
# Create a new migration using Docker
docker run -it guilhermewebdev/migrator:latest migrate new migration_name
```

This will execute the specified Migrator command inside a Docker container based on the provided image.

Note: Ensure that you have Docker installed on your machine. You can find instructions for installing Docker [here](https://docs.docker.com/engine/install/).

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
migrate new <migration_name>

# Execute the next migration
migrate up

# Rollback the last migration
migrate down

# Unlock migrations
migrate unlock

# Perform missing migrations
migrate latest
```

## Environment Variable Configuration

You can apply configurations to Migrator using environment variables. The following variables are supported:

- `DB_DSN`: Database connection string.
- `DB_DRIVER`: Database driver (mysql, postgres, sqlserver, sqlite, sqlite3, or oracle).
- `MIGRATIONS_DIR`: Select the migrations directory (default: "./migrations").
- `MIGRATIONS_TABLE`: Migrations table name (default: "migrations").

To set these variables, you can use your shell's syntax. For example, in Bash:

```bash
export DB_DSN="your_database_connection_string"
export DB_DRIVER="your_database_driver"
export MIGRATIONS_DIR="your_migrations_directory"
export MIGRATIONS_TABLE="your_migrations_table_name"
```

## Configuration File

By default, Migrate looks for a configuration file named "migrator.yml" for global settings. You can specify an alternative configuration file using the `--conf-file` option.

The `migrator.yml` file is used to configure settings for the Migrator command-line tool. Below is an example of the configuration file syntax along with explanations of each parameter:

```yaml
# Example migrator.yml Configuration File

# Directory where migration files are stored
migrations_dir: ./migrations

# Name of the table to track migrations in the database
migrations_table_name: migrations

# Database connection string (DSN)
db_dsn: "postgres://user:pass@postgres:5432/test?sslmode=disable"

# Database driver (Supported drivers: mysql, postgres, sqlserver, sqlite, sqlite3, oracle)
db_driver: postgres
```

**Explanation:**

1. `migrations_dir`: Specifies the directory where migration files are located. In the example, migrations are expected to be in the `./migrations` directory. You can customize this path according to your project structure.

2. `migrations_table_name`: Defines the name of the table used to track migrations in the database. The default is set to "migrations," but you can modify it based on your preferences.

3. `db_dsn`: Represents the Database Source Name (DSN), which contains information about the database connection. In the example, a PostgreSQL database connection string is provided. Update this with the appropriate credentials and connection details for your database.

4. `db_driver`: Specifies the database driver to be used (e.g., mysql, postgres, sqlserver, sqlite, sqlite3, oracle). In the example, the driver is set to "postgres." Choose the appropriate driver based on your database system.

Ensure that the information in the `migrator.yml` file accurately reflects your database setup. You can customize these parameters to suit your project's requirements. If needed, refer to the [Global Options](#global-options) section in the README for additional options that can be specified when running Migrate commands.

---

For additional information on each command and its options, use:

```shell
migrate [command] --help
```

Thank you for using Migrate! If you have any questions or feedback, please contact Guilherme Isaías at <guilherme@guilhermeweb.dev>.