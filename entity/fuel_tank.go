package entity

import "github.com/google/uuid"

type FuelTank struct {
	FromFuel uuid.UUID
	ToTank   uuid.UUID
	quantity float64
}
