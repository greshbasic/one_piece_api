package models

type Location struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Affiliation string `json:"affiliation,omitempty"`
}
