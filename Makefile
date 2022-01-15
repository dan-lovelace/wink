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
migrateup: validate-db-connection
	$(info Migrating up)
	@migrate -path $(MIGRATIONS_DIRECTORY) -database $(DB_CONNECTION_STRING) -verbose up $(count)

.PHONY: migratedown
migratedown: validate-db-connection
	$(info Migrating down)
	@migrate -path $(MIGRATIONS_DIRECTORY) -database $(DB_CONNECTION_STRING) -verbose down $(count)

validate-db-connection:
ifndef DB_CONNECTION_STRING
	$(error Missing DB_CONNECTION_STRING)
endif