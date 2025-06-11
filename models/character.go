package models

type HakiType string

const (
	Armament    HakiType = "armament"
	Observation HakiType = "observation"
	Conquerors  HakiType = "conquerors"
)

type Status string

const (
	Alive   Status = "alive"
	Dead    Status = "dead"
	Missing Status = "missing"
	Unknown Status = "unknown"
)

type Character struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Age             *int        `json:"age,omitempty"`
	Status          *Status     `json:"status,omitempty"`
	Race            *string     `json:"race,omitempty"`
	Origin          *string     `json:"origin,omitempty"`
	FirstAppearance *string     `json:"first_appearance,omitempty"`
	Bounty          *string     `json:"bounty,omitempty"`
	Affiliation     *string     `json:"affiliation,omitempty"`
	Haki            []HakiType  `json:"haki,omitempty"`
	DevilFruit      *DevilFruit `json:"devil_fruit,omitempty"`
}
