postgres:
	docker run --name simplebank -p 5432:5432 -e POSTGRES_USER=postgres -E POSTGRES_PASSWORD=61926114 -d postgres

createdb:
	docker exec -it simplebank createdb --username=postgres --owner=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:61926114@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:61926114@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose down

sqlcinit:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc init

sqlcgenerate:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test: 
	go test -v -cover ./...

.PHONY: postgres createdb migrateup migratedown