.PHONY : help build fmt lint gotest test cover shell image clean
.SILENT:
.DEFAULT_GOAL := run

build:
	go build -o ./.bin/app ./cmd/main.go

run: build

migrate:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

clean:
	go clean
	rm -r ./.bin