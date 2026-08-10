include .env
export
.PHONY: service-run
service-run:

	@go run main.go

 service-deploy:
	docker compose up -d  application 
	#нужно писать название сервиса а не контейнера

service-undeploy:
	docker compose down 


migrate-up:
	migrate -path migrations -database ${CONNECTION} up

migrate-down: 
	migrate -path migrations -database ${CONNECTION} down
