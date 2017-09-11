package models

import (
	"time"

	"gopkg.in/pg.v5"
)

type FetchedChestsUpdater interface {
	UpdateFetchedChests(int, *pg.DB) (int, time.Time, error)
}

type User struct {
	Id                  int64     `json:"id"`
	Identifier          string    `json:"identifier"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	LastFetchedChestsAt time.Time `json:"lastFetchedChestsAt"`
	ChestCount          int       `json:"chestCount"`
}

func (user *User) DecrementChestCount(db *pg.DB) error {
	user.ChestCount -= 1
	return user.Update(db)
}

func (user *User) UpdateFetchedChests(chests int, db *pg.DB) (int, time.Time, error) {
	var err error

	if user.LastFetchedChestsAt.IsZero() || user.canFetchChests() {
		user.ChestCount += chests
		user.LastFetchedChestsAt = time.Now()

		_, err = db.Model(&user).
			Set("chest_count = ?chest_count, last_fetched_chests_at = ?last_fetched_chests_at").
			Returning("*").
			Update()
	}

	return user.ChestCount, user.LastFetchedChestsAt, err
}

func (user *User) Persist(db *pg.DB) error {
	currentTime := time.Now()

	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	err := db.Insert(&user)
	return err
}

func (user *User) Update(db *pg.DB) error {
	currentTime := time.Now()

	user.UpdatedAt = currentTime

	err := db.Update(&user)
	return err
}

func (user *User) canFetchChests() bool {
	return (int(time.Since(user.LastFetchedChestsAt) / time.Hour)) >= 24
}
