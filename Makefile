build:
	docker-compose build app

run:
	docker-compose up app

run_test:
	go test ./... -cover

DB='postgres://postgres:1234@0.0.0.0:5432/property?sslmode=disable'

migrate_up:
	migrate -path ./scripts -database $(DB) up

migrate_down:
	migrate -path ./scripts -database $(DB) down