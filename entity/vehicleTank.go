package entity

import (
	"github.com/google/uuid"
	"time"
)

type VehicleTank struct {
	fromTank   uuid.UUID
	toVehicle  uuid.UUID
	quantity   float64
	odometer   float64
	supplyDate time.Time
}
