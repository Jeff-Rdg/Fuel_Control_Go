package entity

import "github.com/google/uuid"

type Tank struct {
	ID       uuid.UUID
	Capacity float64
	Quantity float64
}
