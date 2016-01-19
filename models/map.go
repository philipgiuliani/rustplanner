package models

// Map represents a snapshot of the world
// swagger:response map
type Map struct {
	ID   int64 `json:"id" db:"id"`
	Seed int32 `json:"seed" db:"seed"`
	Size int16 `json:"size" db:"size"`
}
