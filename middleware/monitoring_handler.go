package middleware

import (
	"github.com/newrelic/go-agent"
	"net/http"
)

type MonitoringHandler struct {
	app newrelic.Application
}

func NewMonitoringHandler(newRelicKey string) *MonitoringHandler {
	config := newrelic.NewConfig("Step Warrior", newRelicKey)
	app, err := newrelic.NewApplication(config)

	if err != nil {
		return nil
	} else {
		return &MonitoringHandler{app: app}
	}
}

func (handler *MonitoringHandler) Monitor(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		txn := handler.app.StartTransaction(r.URL.Path, w, r)
		defer txn.End()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
