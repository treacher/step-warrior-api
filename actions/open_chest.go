package actions

import (
	"github.com/treacher/step-warrior-api/models"
	"github.com/treacher/step-warrior-api/roller"
	"gopkg.in/pg.v5"
)

const itemCount int = 5

type ChestOpener interface {
	Open(user *models.User) *models.Chest
}

type OpenChest struct {
	db     *pg.DB
	roller *roller.Roller
}

func NewChestOpener(db *pg.DB) *OpenChest {
	return &OpenChest{
		db,
		&roller.Roller{},
	}
}

func (co *OpenChest) Open(user *models.User) *models.Chest {
	filledChest := models.Chest{}

	if user.ChestCount > 0 {
		for filledChest.Count() < itemCount {
			itemObject := co.roller.RollForItemObject()
			itemObject = itemObject.GetRandom(user.Id, co.roller.RollForRarity(), co.db)

			if itemObject != nil {
				userItemObject, _ := itemObject.GenerateUserItem(user.Id, co.db)
				filledChest = userItemObject.AppendToChest(&filledChest)
			}
		}

		_ = user.DecrementChestCount(co.db)
	}

	return &filledChest
}
