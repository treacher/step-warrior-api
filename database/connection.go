package database

import (
	"fmt"
	"gopkg.in/pg.v5"

	"github.com/treacher/step-warrior-api/environment"
)

func NewDatabaseConnection(environment string) (db *pg.DB) {
	connection := pg.Connect(options(environment))

	var n int

	_, err := connection.QueryOne(pg.Scan(&n), "SELECT 1")

	panicWhenError(err)

	return connection
}

func options(environmentLabel string) *pg.Options {
	var url string

	if environmentLabel == "test" {
		url = localUrl("step-warrior-api-test")
	} else {
		url = localUrl("step-warrior-api")
	}

	url = environment.Get("DATABASE_URL", url)

	options, err := pg.ParseURL(url)

	panicWhenError(err)

	return options
}

func panicWhenError(err error) {
	if err != nil {
		panic(err)
	}
}

func localUrl(database string) string {
	return fmt.Sprintf("postgres://localhost:5432/%s?sslmode=disable", database)
}
