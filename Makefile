run:
	go run cmd/main.go

tidy:
	go mod tidy

migration_create:
	migrate create -ext sql -dir ./Schemas -seq tables_1

migrate_up:
	migrate -path ./Schemas -database "postgresql://postgres:123456@localhost:5432/task?sslmode=disable" -verbose up

migrate_down:
	migrate -path ./Schemas -database "postgresql://postgres:123456@localhost:5432/ecommerce?sslmode=disable" -verbose down

.PHONY: run tidy migration_create migrate_up migrate_down
