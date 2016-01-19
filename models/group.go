package models

// Group represents a team where users can join
// swagger:response group
type Group struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Map      Map    `json:"map"`
}

// Groups represents a collection of {Group}
// swagger:response groups
type Groups []Group
