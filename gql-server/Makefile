build:
	@go build -o bin/gql-server

run: build
	./bin/gql-server

migrate:
	sh ./scripts/migrate.sh

generate:
	sh ./scripts/generate.sh

test:
	go test -v ./...