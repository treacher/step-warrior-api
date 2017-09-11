package models

import (
	"time"

	"gopkg.in/pg.v5"
)

type UserEquipment struct {
	TableName   struct{}                 `sql:"user_equipment"`
	Id          int64                    `json:"id"`
	EquipmentId int64                    `json:"equipmentId"`
	UserId      int64                    `json:"userId"`
	Attributes  []UserEquipmentAttribute `json:"attributes"`
	CreatedAt   time.Time                `json:"createdAt"`
	UpdatedAt   time.Time                `json:"updatedAt"`
	Equipment   *Equipment
}

type UserEquipmentAttribute struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

func GetUserEquipment(db *pg.DB, userId int64) ([]UserEquipment, error) {
	var userEquipment []UserEquipment

	err := db.Model(&userEquipment).
		Column("user_equipment.*", "Equipment").
		Where("user_id = ?", userId).
		Select()

	return userEquipment, err
}

func (userEquipment *UserEquipment) AppendToChest(chest *Chest) Chest {
	chest.UserEquipment = append(chest.UserEquipment, userEquipment)
	return *chest
}

func (userEquipment *UserEquipment) Persist(db *pg.DB) error {
	var currentTime = time.Now()

	userEquipment.CreatedAt = currentTime
	userEquipment.UpdatedAt = currentTime

	err := db.Insert(&userEquipment)
	return err
}
