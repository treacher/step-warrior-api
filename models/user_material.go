package models

import (
	"time"

	"gopkg.in/pg.v5"
)

type UserMaterial struct {
	Id         int64     `json:"id"`
	MaterialId int64     `json:"materialId"`
	UserId     int64     `json:"userId"`
	Quantity   int64     `json:"quantity"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Material   *Material
}

func GetUserMaterials(db *pg.DB, userId int64) ([]UserMaterial, error) {
	var userMaterials []UserMaterial

	err := db.Model(&userMaterials).
		Column("user_material.*", "Material").
		Where("user_id = ?", userId).
		Select()

	return userMaterials, err
}

func (userMaterial *UserMaterial) Update(db *pg.DB) error {
	var currentTime = time.Now()

	userMaterial.UpdatedAt = currentTime

	err := db.Update(&userMaterial)
	return err
}

func (userMaterial *UserMaterial) Persist(db *pg.DB) error {
	var currentTime = time.Now()

	userMaterial.CreatedAt = currentTime
	userMaterial.UpdatedAt = currentTime

	err := db.Insert(&userMaterial)
	return err
}

func CreateOrUpdateCountForUserMaterial(userMaterial *UserMaterial, db *pg.DB) error {
	db.Model(userMaterial).
		Where("user_id = ?", userMaterial.UserId).
		Where("material_id = ?", userMaterial.MaterialId).
		Limit(1).
		Select()

	userMaterial.Quantity = userMaterial.Quantity + 1

	if userMaterial.Id == 0 {
		return userMaterial.Persist(db)
	}

	return userMaterial.Update(db)
}

func (userMaterial *UserMaterial) AppendToChest(chest *Chest) Chest {
	chest.UserMaterials = append(chest.UserMaterials, userMaterial)
	return *chest
}
