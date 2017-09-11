package actions

import (
	"github.com/treacher/step-warrior-api/models"
	"gopkg.in/pg.v5"

	"time"
)

type ChestFetcher interface {
	Fetch(user models.FetchedChestsUpdater, steps int) (*FetchChestResponse, error)
}

type FetchChest struct {
	db *pg.DB
}

type FetchChestResponse struct {
	ChestCount          int       `json:"chestCount"`
	LastFetchedChestsAt time.Time `json:"lastFetchedChestsAt"`
}

func NewChestFetcher(db *pg.DB) *FetchChest {
	return &FetchChest{db: db}
}

func (cf *FetchChest) Fetch(user models.FetchedChestsUpdater, steps int) (*FetchChestResponse, error) {
	chestCount, lastFetchedChestsAt, err :=
		user.UpdateFetchedChests(cf.calculateAmountOfChests(steps), cf.db)

	return &FetchChestResponse{
		ChestCount:          chestCount,
		LastFetchedChestsAt: lastFetchedChestsAt,
	}, err
}

func (cf *FetchChest) calculateAmountOfChests(steps int) int {
	if steps >= 20000 {
		steps = 20000
	}
	return int(float32(steps)/float32(1500) + 0.5)
}
