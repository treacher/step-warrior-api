package models

import (
	"crypto/rand"
	"math/big"
	"time"

	"gopkg.in/pg.v5"
)

type Equipment struct {
	Id         int64                `json:"id"`
	Name       string               `json:"name"`
	Rarity     string               `json:"rarity"`
	Attributes []EquipmentAttribute `json:"attributes"`
	SlotType   string               `json:"slotType"`
	CreatedAt  time.Time            `json:"createdAt"`
	UpdatedAt  time.Time            `json:"updatedAt"`
}

type EquipmentAttribute struct {
	Type string `json:"type"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

func (equipment *Equipment) GetRandom(userId int64, rarity string, db *pg.DB) Item {
	var randomEquipment Equipment

	db.Model(&randomEquipment).
		Where("rarity = ?", rarity).
		OrderExpr("Random()").
		Limit(1).Select()

	if randomEquipment.Id == 0 {
		return nil
	}

	return &randomEquipment
}

func (equipment *Equipment) GenerateUserItem(userId int64, db *pg.DB) (UserItem, error) {
	var userEquipmentAttributes []UserEquipmentAttribute

	for _, attribute := range equipment.Attributes {
		attrValue := equipment.getRandomAttribute(attribute.Min, attribute.Max)
		userEquipmentAttr := UserEquipmentAttribute{Type: attribute.Type, Value: attrValue}
		userEquipmentAttributes = append(userEquipmentAttributes, userEquipmentAttr)
	}

	userEquipment := &UserEquipment{
		Attributes:  userEquipmentAttributes,
		EquipmentId: equipment.Id,
		UserId:      userId,
	}

	err := userEquipment.Persist(db)

	userEquipment.Equipment = equipment

	return userEquipment, err
}

func (equipment *Equipment) getRandomAttribute(min int, max int) int64 {
	number, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))

	if err != nil {
		println(err)
	}

	return number.Int64() + int64(min)
}
