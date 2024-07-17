run:
	go run main.go

tidy:
	go mod tidy

migration_author_crt:
	migrate create -ext sql -dir ./Schemas/author/ -seq author_1

migration_book_crt:
	migrate create -ext sql -dir ./Schemas/book/ -seq book_1
