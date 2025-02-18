schema:
	migrate create -ext sql -dir db/migration -seq $(name)
	
migrateup:
	migrate -database postgres://root:root@localhost:5432/nanachat?sslmode=disable -path db/migration up

migratedown:
	migrate -database postgres://root:root@localhost:5432/nanachat?sslmode=disable -path db/migration down

postgres:
	docker run --name nanachat-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:17.3-alpine

sqlc:
	sqlc generate

test:
	go test -v -cover ./... 

.PHONY: schema migrateup migratedown postgres sqlc test