env:
	docker-compose up redis

run:
	go run ./bin/main.go

start:
	docker-compose build
	docker-compose up api

stop:
	docker-compose stop

build:
	go build .