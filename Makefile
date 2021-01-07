APP=app

build:
	docker-compose build $(app)

run:
	docker-compose up $(app)

run_test:
	go test ./... -cover

SCHEMA=./scripts
DB='postgres://postgres:1234@0.0.0.0:5436/postgres?sslmode=disable'

migrate_up:
	migrate -path $(SCHEMA) -database $(DB) up

migrate_down:
	migrate -path $(SCHEMA) -database $(DB) down