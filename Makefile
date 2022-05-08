postgres:
	docker run --name golang_bank_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

dbstart:
	docker start ce1c810f330c

dbstop:
	docker stop ce1c810f330c

createdb:
	docker exec -it golang_bank_postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it golang_bank_postgres dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres dbstart dbstop createdb dropdb migrateup migratedown sqlc test