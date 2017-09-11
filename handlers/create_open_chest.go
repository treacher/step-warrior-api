package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/treacher/step-warrior-api/actions"
	"github.com/treacher/step-warrior-api/models"
	"github.com/treacher/step-warrior-api/presenters"
	"gopkg.in/pg.v5"
)

type CreateOpenChest struct {
	chestOpener actions.ChestOpener
	db          *pg.DB
}

func NewCreateOpenChest(chestOpener actions.ChestOpener) *CreateOpenChest {
	return &CreateOpenChest{
		chestOpener: chestOpener,
	}
}

func (cci *CreateOpenChest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("User").(models.User)

	chestItems := cci.chestOpener.Open(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(presenters.NewChest(chestItems, &user))
}
