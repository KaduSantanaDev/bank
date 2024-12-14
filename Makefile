include .env
export $(shell sed 's/=.*//' .env)

CONN_STRING := postgresql://${AWS_DB_USER}:${AWS_DB_PASSWORD}@${AWS_DB_HOST}:${AWS_DB_PORT}/${AWS_DB_NAME}

postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -t postgres12 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres12 dropdb bank

migrateup:
	migrate -path db/migration -database "${CONN_STRING}" -verbose up

migrateup1:
	migrate -path db/migration -database "${CONN_STRING}" -verbose up 1

migratedown:
	migrate -path db/migration -database "${CONN_STRING}" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres, createdb, migrateup, migratedown, sqlc, test, server, migrateup1, migratedown1