# wink

A command line utility for tracking time written in Go

## Getting Started

TODO

# Development

## Data Migrations

Migrations are handled using [golang-migrate](https://github.com/golang-migrate/migrate). You will need to install the command line utility [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to run the necessary commands.

#### Commands

- New migration: `make newmigration name=<migration_name>` - Replace `<migration_name>` with a brief description of the changes. A good migration name would be something like `add_active_column_to_user_table`.
- Migrate up: `make migrateup` - Runs ALL up migrations.
- Migrate down: `make migratedown` - Runs ALL down migrations.
