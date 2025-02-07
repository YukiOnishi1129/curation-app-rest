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
	migrate create -ext sql -dir db/migrations -seq $(name)

migration-up:
	migrate -path db/migrations -database "postgres://user:pass@localhost:5432/CURATION_APP_REST_DB?sslmode=disable" up

migration-rollback:
	migrate -path db/migrations -database "postgres://user:pass@localhost:5432/CURATION_APP_REST_DB?sslmode=disable" down