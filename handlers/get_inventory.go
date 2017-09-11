package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/treacher/step-warrior-api/actions"
	"github.com/treacher/step-warrior-api/middleware"
	"github.com/treacher/step-warrior-api/models"
	"github.com/treacher/step-warrior-api/presenters"
)

type GetInventory struct {
	inventoryGetter *actions.InventoryGetter
}

func NewGetInventory(inventoryGetter *actions.InventoryGetter) *GetInventory {
	return &GetInventory{
		inventoryGetter: inventoryGetter,
	}
}

func (gi *GetInventory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("User").(models.User)
	inventory, err := gi.inventoryGetter.GetInventory(&user)

	if err != nil {
		switch err {
		case actions.InventoryFetchError:
			// Return some error
			// Error(r.Context(), w, 404, "not_found", "biller not found")
		default:
			middleware.InternalServerError(w, r)
		}
		return
	}

	inventoryResponse := presenters.GetInventoryResponse(inventory)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventoryResponse)
}
