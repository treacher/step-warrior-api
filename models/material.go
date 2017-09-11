package models

import (
	"gopkg.in/pg.v5"
	"time"
)

type Material struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Rarity    string    `json:"rarity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (material *Material) GenerateUserItem(userId int64, db *pg.DB) (UserItem, error) {
	userMaterial := &UserMaterial{UserId: userId, MaterialId: material.Id}
	err := CreateOrUpdateCountForUserMaterial(userMaterial, db)
	userMaterial.Material = material

	return userMaterial, err
}

func (material *Material) GetRandom(userId int64, rarity string, db *pg.DB) Item {
	var randomMaterial Material

	db.Model(&randomMaterial).
		Where("rarity = ?", rarity).
		OrderExpr("Random()").
		Limit(1).Select()

	if randomMaterial.Id == 0 {
		return nil
	}

	return &randomMaterial
}

func (material *Material) Update(db *pg.DB) error {
	var currentTime = time.Now()

	material.UpdatedAt = currentTime

	err := db.Update(&material)
	return err
}
