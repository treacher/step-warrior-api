package models

import (
	"gopkg.in/pg.v5"
	"time"
)

type Plan struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Rarity      string    `json:"rarity"`
	EquipmentId int64     `json:"equipmentId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (plan *Plan) GenerateUserItem(userId int64, db *pg.DB) (UserItem, error) {
	userPlan := &UserPlan{UserId: userId, PlanId: plan.Id}
	err := userPlan.Persist(db)
	userPlan.Plan = plan

	return userPlan, err
}

func (plan *Plan) GetRandom(userId int64, rarity string, db *pg.DB) Item {
	var randomPlan Plan
	var ownedPlanIds []int

	db.Model(&UserPlan{}).
		ColumnExpr("array_agg(user_plan.plan_id)").
		Where("user_id = ?", userId).
		Select(pg.Array(ownedPlanIds))

	db.Model(&randomPlan).
		Where("rarity = ?", rarity).
		Where("plan_id NOT IN (?)", ownedPlanIds).
		OrderExpr("Random()").Limit(1).Select()

	if randomPlan.Id == 0 {
		return nil
	}

	return &randomPlan
}
