.PHONY: install format test db_migrate test_migrate coverage build

install:
	glide install

format:
	go fmt $(shell glide novendor)

test: test_migrate
	go test -v $(shell glide novendor)

db_migrate:
	goose -path database up

test_migrate:
	goose -env test -path database up

coverage: install format
	go test -cover $(shell glide novendor)

build: install format
	go build -v -o ./build/step-warrior-api
