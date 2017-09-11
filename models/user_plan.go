package models

import (
	"time"

	"gopkg.in/pg.v5"
)

type UserPlan struct {
	Id        int64     `json:"id"`
	PlanId    int64     `json:"planId"`
	UserId    int64     `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Plan      *Plan
}

func GetUserPlans(db *pg.DB, userId int64) ([]UserPlan, error) {
	var userPlans []UserPlan

	err := db.Model(&userPlans).
		Column("user_plan.*", "Plan").
		Where("user_id = ?", userId).
		Select()

	return userPlans, err
}

func (userPlan *UserPlan) AppendToChest(chest *Chest) Chest {
	chest.UserPlans = append(chest.UserPlans, userPlan)
	return *chest
}

func (userPlan *UserPlan) Update(db *pg.DB) error {
	var currentTime = time.Now()
	userPlan.UpdatedAt = currentTime

	return db.Update(&userPlan)
}

func (userPlan *UserPlan) Persist(db *pg.DB) error {
	var currentTime = time.Now()

	userPlan.CreatedAt = currentTime
	userPlan.UpdatedAt = currentTime

	return db.Insert(&userPlan)
}
