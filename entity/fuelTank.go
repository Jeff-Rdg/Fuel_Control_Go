package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	FuelTankQuantityError = errors.New("invalid quantity to supply")
	FuelTankDateError     = errors.New("invalid date to supply")
)

type FuelTank struct {
	fromFuel   uuid.UUID
	toTank     uuid.UUID
	quantity   float64
	supplyDate time.Time
}

func NewFuelTank(fromFuel, toTank uuid.UUID, quantity float64, date time.Time) (FuelTank, error) {
	if quantity <= 0 {
		return FuelTank{}, FuelTankQuantityError
	}
	if date.After(time.Now()) {
		return FuelTank{}, FuelTankDateError
	}
	return FuelTank{
		fromFuel:   fromFuel,
		toTank:     toTank,
		quantity:   quantity,
		supplyDate: date,
	}, nil
}
