package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/treacher/step-warrior-api/actions"
	"github.com/treacher/step-warrior-api/models"
)

type CreateChestsInput struct {
	Steps int `json:"steps"`
}

type CreateChests struct {
	chestFetcher actions.ChestFetcher
}

func NewCreateChests(chestFetcher actions.ChestFetcher) *CreateChests {
	return &CreateChests{chestFetcher: chestFetcher}
}

func (cc *CreateChests) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	createChestsInput := &CreateChestsInput{}

	user := r.Context().Value("User").(models.User)

	_ = json.NewDecoder(r.Body).Decode(createChestsInput)

	chestFetcher, _ := cc.chestFetcher.Fetch(&user, createChestsInput.Steps)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(chestFetcher)
}
