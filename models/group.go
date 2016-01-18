package models

type Group struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
