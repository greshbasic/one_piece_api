package models

type DevilFruit struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Awakened bool   `json:"awakened"`
	UserID   *int64 `json:"user_id,omitempty"`
}
