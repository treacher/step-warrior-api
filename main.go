package main

import (
	"log"
	"net/http"

	"github.com/treacher/step-warrior-api/database"
	"github.com/treacher/step-warrior-api/environment"
	"github.com/treacher/step-warrior-api/router"
)

func main() {
	port := environment.Get("PORT", "8080")
	router := router.NewRouter(database.NewDatabaseConnection("production"))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
