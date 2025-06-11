package models

type DevilFruitType string

const (
	Paramecia DevilFruitType = "paramecia"
	Zoan      DevilFruitType = "zoan"
	Ancient   DevilFruitType = "ancient zoan"
	Mythical  DevilFruitType = "mythical zoan"
	Logia     DevilFruitType = "logia"
)

type DevilFruit struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Awakened bool    `json:"awakened"`
	UserID   *string `json:"user_id,omitempty"`
}
