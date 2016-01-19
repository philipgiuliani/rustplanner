package models

// Group represents a team where users can join
// swagger:response group
type Group struct {
	ID       int64  `json:"id" db:"id,omitempty"`
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
	Map      *Map   `json:"map,omitempty" db:"-"`
}

// Groups represents a collection of {Group}
// swagger:response groups
type Groups []Group
