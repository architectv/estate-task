build:
	docker-compose build app

run:
	docker-compose up app

run_test:
	go test ./... -cover

migrate_up:
	migrate -path ./scripts -database 'postgres://postgres:1234@0.0.0.0:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./scripts -database 'postgres://postgres:1234@0.0.0.0:5436/postgres?sslmode=disable' down