package presenters

import (
	"github.com/treacher/step-warrior-api/models"
)

type Chest struct {
	ChestCount int              `json:"chestCount"`
	Materials  []*UserMaterial  `json:"materials"`
	Equipment  []*UserEquipment `json:"equipment"`
	Plans      []*UserPlan      `json:"plans"`
}

func NewChest(chest *models.Chest, user *models.User) *Chest {
	var chestPresenter Chest

	chestPresenter.ChestCount = user.ChestCount

	for _, userEquipment := range chest.UserEquipment {
		chestPresenter.Equipment = append(chestPresenter.Equipment, NewUserEquipment(userEquipment))
	}

	for _, userMaterial := range chest.UserMaterials {
		chestPresenter.Materials = append(chestPresenter.Materials, NewUserMaterial(userMaterial))
	}

	for _, userPlan := range chest.UserPlans {
		chestPresenter.Plans = append(chestPresenter.Plans, NewUserPlan(userPlan))
	}

	return &chestPresenter
}

type UserEquipment struct {
	Name       string                          `json:"name"`
	Rarity     string                          `json:"rarity"`
	Attributes []models.UserEquipmentAttribute `json:"attributes"`
}

func NewUserEquipment(userEquipment *models.UserEquipment) *UserEquipment {
	equipment := userEquipment.Equipment

	return &UserEquipment{
		Name:       equipment.Name,
		Rarity:     equipment.Rarity,
		Attributes: userEquipment.Attributes,
	}
}

type UserMaterial struct {
	Name   string `json:"name"`
	Rarity string `json:"rarity"`
}

func NewUserMaterial(userMaterial *models.UserMaterial) *UserMaterial {
	material := userMaterial.Material

	return &UserMaterial{
		Name:   material.Name,
		Rarity: material.Rarity,
	}
}

type UserPlan struct {
	Name   string `json:"name"`
	Rarity string `json:"rarity"`
}

func NewUserPlan(userPlan *models.UserPlan) *UserPlan {
	plan := userPlan.Plan

	return &UserPlan{
		Name:   plan.Name,
		Rarity: plan.Rarity,
	}
}
