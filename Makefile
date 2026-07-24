include .env
export

service-run:
export connection="postgres://postgres:1234@localhost:5432/postgres" && \
go run main.go




migrate-up:
	migrate -path migrations -database ${CONNECTION} up

migrate-down: 
	migrate -path migrations -database ${CONNECTION} down
