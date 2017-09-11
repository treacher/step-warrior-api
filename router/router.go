package router

import (
	"net/http"
	"os"

	"github.com/pressly/chi"
	chiMiddleware "github.com/pressly/chi/middleware"
	"github.com/treacher/step-warrior-api/actions"
	"github.com/treacher/step-warrior-api/handlers"
	"github.com/treacher/step-warrior-api/middleware"
	"gopkg.in/pg.v5"
)

func NewRouter(db *pg.DB) http.Handler {
	r := chi.NewRouter()

	monitoringHandler := middleware.NewMonitoringHandler(os.Getenv("NEW_RELIC_LICENSE_KEY"))
	authenticationHandler := middleware.NewOAuthHandler(db)

	r.Use(chiMiddleware.Heartbeat("/"))
	r.Use(chiMiddleware.Heartbeat("/healthz"))
	r.Use(authenticationHandler.OAuthHandler)

	if monitoringHandler != nil {
		r.Use(monitoringHandler.Monitor)
	}

	chestOpener := actions.NewChestOpener(db)
	r.Post("/chest_items", handlers.NewCreateOpenChest(chestOpener).ServeHTTP)

	chestFetcher := actions.NewChestFetcher(db)
	r.Post("/chests", handlers.NewCreateChests(chestFetcher).ServeHTTP)

	inventoryGetter := actions.NewInventoryGetter(db)
	r.Get("/inventory", handlers.NewGetInventory(inventoryGetter).ServeHTTP)

	return r
}
