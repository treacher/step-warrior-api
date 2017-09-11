package actions

import (
	"errors"

	"github.com/treacher/step-warrior-api/models"
	"gopkg.in/pg.v5"
)

var InventoryFetchError error = errors.New("Error while retreiving inventory")

func NewInventoryGetter(db *pg.DB) *InventoryGetter {
	return &InventoryGetter{db: db}
}

type InventoryGetter struct {
	db *pg.DB
}

type Inventory struct {
	UserMaterials []models.UserMaterial  `json:"user_materials"`
	UserEquipment []models.UserEquipment `json:"user_equipment"`
	UserPlans     []models.UserPlan      `json:"user_plans"`
}

func (ig *InventoryGetter) GetInventory(user *models.User) (*Inventory, error) {

	userEquipment, equipmentFetchErr := models.GetUserEquipment(ig.db, user.Id)

	if equipmentFetchErr != nil {
		return nil, InventoryFetchError
	}

	userMaterials, materialsFetchErr := models.GetUserMaterials(ig.db, user.Id)

	if materialsFetchErr != nil {
		return nil, InventoryFetchError
	}

	userPlans, plansFetchErr := models.GetUserPlans(ig.db, user.Id)

	if plansFetchErr != nil {
		return nil, InventoryFetchError
	}

	// TODO duplicates OpenedChest struct
	inventory := &Inventory{
		UserMaterials: userMaterials,
		UserEquipment: userEquipment,
		UserPlans:     userPlans,
	}
	return inventory, nil
}
