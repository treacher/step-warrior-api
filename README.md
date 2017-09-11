# Game API

## Development

### Run database migrations

* Run `go get bitbucket.org/liamstask/goose/cmd/goose` to install goose
* Run `goose up`

### Run the application

* Install [glide](https://github.com/Masterminds/glide) by following the instructions on the website.
* Run `make build`
* Run `./build/game-api`

### Run the tests
If you're runing for the first time run the following:
`createdb step-warrior-api-test`
`psql -Upostgres`
`\c step-warrior-api-test`
`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

Otherwise run:
`make tests`

#### Endpoints

##### Open a chest

This endpoint requires you to have atleast one item in the database or it will get stuck in an infinite loop as it wont stop until it has 5 items.

POST http://localhost:8080/chest_items

Payload:
```JSON
{
  "materials" : [],
  "equipment" : [],
  "plans" :[]
}
```
