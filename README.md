# wink

A command line utility for tracking time written in Go

## Getting Started

TODO

## Data Migrations

Migrations are handled using [golang-migrate](https://github.com/golang-migrate/migrate). You will need to install the command line utility [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to run the necessary commands. You need to set the variable `DB_CONNECTION_STRING` before running migrations.

#### New migration
```
make newmigration name=my_new_migration
```

#### Migrate UP

```
DB_CONNECTION_STRING=<my_connection_string> make migrateup
```

#### Migrate DOWN with optional `count` parameter (also works for `migrateup`)
```
DB_CONNECTION_STRING=<my_connection_string> make migratedown count=1
```
