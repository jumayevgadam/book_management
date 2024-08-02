ifeq ($(POSTGRES_SETUP),)
	POSTGRES_SETUP := user=postgres password=123456 dbname=library host=localhost port=5432 sslmode=disable
endif

MIGRATION_FOLDER=$(CURDIR)/schemas


.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down

run:
	go run cmd/main.go

tidy:
	go mod tidy

swag:
	swag init -g cmd/main.go	

.PHONY: run tidy migration_create migrate_up migrate_down migration_fix swag	
