package models

type UserItem interface {
	AppendToChest(chest *Chest) Chest
}
