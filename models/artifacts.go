package models

type Artifact struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Origin      *string `json:"origin,omitempty"`
}
