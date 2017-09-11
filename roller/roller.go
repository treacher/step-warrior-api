package roller

import (
	"crypto/rand"
	"math/big"

	"github.com/treacher/step-warrior-api/models"
)

const (
	COMMON     = "common"
	RARE       = "rare"
	SUPER_RARE = "super"
	ULTRA_RARE = "ultra"
	ANCIENT    = "ancient"
)

type Roller struct{}

func (roller *Roller) RollForItemObject() models.Item {
	diceRoll := roller.getRandomNumber(100)

	switch {
	case diceRoll > 60 && diceRoll <= 90:
		return &models.Plan{}
	case diceRoll > 90 && diceRoll <= 100:
		return &models.Equipment{}
	default:
		return &models.Material{}
	}
}

func (roller *Roller) RollForRarity() string {
	diceRoll := roller.getRandomNumber(10001)

	rarity := COMMON // 60%

	switch {
	case diceRoll > 8399 && diceRoll <= 9399: // 10%
		rarity = RARE
	case diceRoll > 9400 && diceRoll <= 9899: // 4.99%
		rarity = SUPER_RARE
	case diceRoll > 9900 && diceRoll <= 9999: // 1%
		rarity = ULTRA_RARE
	case diceRoll == 10000: // 0.001%
		rarity = ANCIENT
	}

	return rarity
}

func (roller *Roller) getRandomNumber(max int64) int64 {
	number, err := rand.Int(rand.Reader, big.NewInt(int64(max)))

	if err != nil {
		println(err)
	}

	return number.Int64()
}
