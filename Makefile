postgrescont:
	docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:14.5-alpine

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/simplebank?sslmode=disable" -verbose down

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simplebank
	
dropdb:
	docker exec -it postgres dropdb --username=postgres simplebank

sqlcgen:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

dbtest:
	go test -v -cover ./...

.PHONY:
	postgrescont createdb dropdb migrateup migratedown dbtest