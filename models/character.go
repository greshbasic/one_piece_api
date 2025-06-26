package models

type HakiType string

type HakiAwakening struct {
	Awakened bool `json:"awakened"`
}

type Haki map[string]HakiAwakening

type Character struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Bounty      *int64      `json:"bounty,omitempty"`
	Haki        *Haki       `json:"haki,omitempty"`
	Affiliation *string     `json:"affiliation,omitempty"`
	Origin      *string     `json:"origin,omitempty"`
	Status      *string     `json:"status,omitempty"`
	Age         *int        `json:"age,omitempty"`
	DevilFruit  *DevilFruit `json:"devil_fruit,omitempty"`
}
