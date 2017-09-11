package models

import "gopkg.in/pg.v5"

type Item interface {
	GetRandom(userId int64, rarity string, db *pg.DB) Item
	GenerateUserItem(userId int64, db *pg.DB) (UserItem, error)
}
