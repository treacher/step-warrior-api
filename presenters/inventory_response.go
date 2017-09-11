package presenters

import (
	"github.com/treacher/step-warrior-api/actions"
	"github.com/treacher/step-warrior-api/models"
)

type InventoryResponse struct {
	Materials []*inventoryMaterial  `json:"materials"`
	Equipment []*inventoryEquipment `json:"equipment"`
	Plans     []*inventoryPlan      `json:"plans"`
}

func GetInventoryResponse(inventory *actions.Inventory) *InventoryResponse {
	var inventoryResponse InventoryResponse

	for _, userEquipment := range inventory.UserEquipment {
		inventoryResponse.Equipment = append(inventoryResponse.Equipment, newInventoryEquipment(userEquipment))
	}

	for _, userMaterial := range inventory.UserMaterials {
		inventoryResponse.Materials = append(inventoryResponse.Materials, newInventoryMaterial(userMaterial))
	}

	for _, userPlan := range inventory.UserPlans {
		inventoryResponse.Plans = append(inventoryResponse.Plans, newInventoryPlan(userPlan))
	}

	return &inventoryResponse
}

type inventoryEquipment struct {
	Name       string                          `json:"name"`
	Rarity     string                          `json:"rarity"`
	Attributes []models.UserEquipmentAttribute `json:"attributes"`
}

func newInventoryEquipment(userEquipment models.UserEquipment) *inventoryEquipment {
	equipment := userEquipment.Equipment

	return &inventoryEquipment{
		Name:       equipment.Name,
		Rarity:     equipment.Rarity,
		Attributes: userEquipment.Attributes,
	}
}

type inventoryMaterial struct {
	Name     string `json:"name"`
	Rarity   string `json:"rarity"`
	Quantity int64  `json:"quantity"`
}

func newInventoryMaterial(userMaterial models.UserMaterial) *inventoryMaterial {
	material := userMaterial.Material

	return &inventoryMaterial{
		Name:     material.Name,
		Rarity:   material.Rarity,
		Quantity: userMaterial.Quantity,
	}
}

type inventoryPlan struct {
	Name   string `json:"name"`
	Rarity string `json:"rarity"`
}

func newInventoryPlan(userPlan models.UserPlan) *inventoryPlan {
	plan := userPlan.Plan

	return &inventoryPlan{
		Name:   plan.Name,
		Rarity: plan.Rarity,
	}
}
