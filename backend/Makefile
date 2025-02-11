include .env.local

empty:
	echo "empty"

lint:
	golangci-lint run ./...

test:
	go test -v ./...

## make test-single name=TestFunctionName
test-single:
	go test -v -run ${name} ./...

create-migration:
	migrate create -ext sql -dir database/migrations -seq $(name)

migration-up:
	migrate -path migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" up

migration-rollback:
	migrate -path migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" down

gen-sqlc:
	sqlc generate