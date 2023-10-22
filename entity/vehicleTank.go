package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	VehicleTankQuantityError = errors.New("invalid quantity to supply")
	VehicleTankDateError     = errors.New("invalid date to supply")
)

type VehicleTank struct {
	fromTank   uuid.UUID
	toVehicle  uuid.UUID
	quantity   float64
	odometer   int
	supplyDate time.Time
}

func NewVehicleTank(fromTank, toVehicle uuid.UUID, quantity float64, odometer int, date time.Time) (VehicleTank, error) {
	if quantity <= 0 {
		return VehicleTank{}, VehicleTankQuantityError
	}
	if odometer < 0 {
		return VehicleTank{}, VehicleOdometerError
	}
	if date.After(time.Now()) {
		return VehicleTank{}, VehicleTankDateError
	}
	return VehicleTank{
		fromTank:   fromTank,
		toVehicle:  toVehicle,
		quantity:   quantity,
		odometer:   odometer,
		supplyDate: date,
	}, nil
}
