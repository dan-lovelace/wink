MIGRATIONS_DIRECTORY="db/migrations"

.PHONY: newmigration
newmigration:
ifndef name
	$(info Invalid usage)
	$(info Expected: make newmigration name=my_new_migration)
	$(error Missing migration name)
endif

	$(info Creating new migration $(name))
	migrate create -ext sql -dir $(MIGRATIONS_DIRECTORY) -seq $(name)

.PHONY: migrateup
migrateup:
	@go run migrate/main.go up

.PHONY: migratedown
migratedown:
	@go run migrate/main.go down

validate-db-connection:
ifndef DB_CONNECTION_STRING
	$(error Missing DB_CONNECTION_STRING)
endif