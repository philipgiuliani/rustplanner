package models

type Map struct {
	ID    int64  `json:"id"`
	Seed  int32  `json:"seed"`
	Size  int16  `json:"size"`
	Image string `json:"-"`
}
