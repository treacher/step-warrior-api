package models

type Chest struct {
	UserMaterials []*UserMaterial  `json:"materials"`
	UserEquipment []*UserEquipment `json:"equipment"`
	UserPlans     []*UserPlan      `json:"plans"`
}

func (chest *Chest) Count() int {
	return len(chest.UserMaterials) + len(chest.UserEquipment) + len(chest.UserPlans)
}
