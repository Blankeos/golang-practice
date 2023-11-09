build:
	@go build -o bin/gql-server

run: build
	./bin/gql-server

migrate:
	source ./scripts/create_migration.sh

generate:
	source ./scripts/generate.sh

test:
	go test -v ./...